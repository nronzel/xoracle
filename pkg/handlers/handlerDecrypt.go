package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	dc "github.com/nronzel/xoracle/pkg/decryption"
	tmpls "github.com/nronzel/xoracle/templates"
	"github.com/nronzel/xoracle/utils"
)

func HandlerDecrypt(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Header().Set("Content-Type", "text-/plain")
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	encodedData := r.FormValue("inputData")

	// Checks if data is Base64 or Hex encoded, and decodes, otherwise just
	// returns the data as is. It's either plaintext, or some other encoding
	// not checked for.
	verifiedData, err := utils.Decode(encodedData)
	if err != nil {
		w.Header().Set("Content-Type", "text/plaintext")
	}

	topKeySizes, err := dc.GuessKeySizes(verifiedData)
	if err != nil {
		w.Header().Set("Content-Type", "text/plaintext")
		http.Error(w, "problem guessing keysizes", http.StatusInternalServerError)
		return
	}

	if len(verifiedData) == 0 || len(topKeySizes) == 0 {
		w.Header().Set("Content-Type", "text/plaintext")
		http.Error(w, "data or key sizes are missing", http.StatusBadRequest)
		return
	}

	results := dc.ProcessKeySizes(topKeySizes, verifiedData)

	// Initialize the HTML response with a container div
	var responseHTML strings.Builder
	responseHTML.WriteString(`<div class="decryption-results">`)

	// Score the decrypted text for each attempted key, return the
	// result that is most likely English text (highest score).
	best := dc.ScoreResults(results)

	responseHTML.WriteString(fmt.Sprintf(
		`<div class="decryption-result">
            <h3>Key Size: %d</h3>
            <h3>Key: %s</h3>
            <p>Decrypted Data:</p>
            <pre>%s</pre>
        </div>`,
		best.KeySize,
		template.HTMLEscapeString(string(best.Key)),
		template.HTMLEscapeString(string(best.DecryptedData)),
	))

	// Close the container div
	responseHTML.WriteString(`</div>`)

	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, responseHTML.String())
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(tmpls.IndexTemplate)
	if err != nil {
		w.Header().Set("Content-Type", "text-/plain")
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.Header().Set("Content-Type", "text-/plain")
		http.Error(w, "problem executing template", http.StatusInternalServerError)
	}
}
