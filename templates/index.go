package templates

const IndexTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Repeating-Key XOR Cracker</title>
    <!-- Include HTMX -->
    <script src="https://unpkg.com/htmx.org@1.9.10" integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC" crossorigin="anonymous"></script>
    <script src="https://unpkg.com/htmx.org/dist/ext/response-targets.js"></script>
    <style>
        * {
            box-sizing: border-box;
        }
        body, html {
            margin: 0;
            padding: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #0e0c17;
            color: #333;
        }
        #container {
            padding: 20px;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 30px;
            width: 100%;
            max-width: 1500px;
            background-color: #25242c;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
            margin: 25px 20px 0px 20px !important;
            min-height: 50vh;
        }
        h1 {
            color: #b886e2;
        }
        form {
            display: flex;
            flex-direction: row;
            align-items: start;
            justify-content: space-around;
            width: 100%;
            gap: 15px;
            color: white;
        }
        .text-container, .output-container {
            flex: 1;
            display: flex;
            flex-direction: column;
        }
        .output-container {
            white-space: nowrap;
            overflow-x: auto;
            overflow-y: auto;
            padding: 10px;
            border: 1px solid #ccc;
            border-radius: 5px;
            background-color: #fafafa;
            margin-top: 18px;
            background-color: #333333;
            color: white;
            max-height: 600px;
        }
        textarea {
            border-radius: 5px;
            border: 1px solid #ccc;
            box-sizing: border-box;
            background-color: #333333 !important;
            color: white;
            margin: 0 !important;
            overflow-x: auto;
        }
        textarea, #output {
            width: 100%;
            padding: 10px;
        }
        textarea {
            height: 300px;
            resize: none;
            background-color: #fff;
        }
        button {
            cursor: pointer;
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            background-color: #b886e2;
            color: white;
            transition: background-color 0.3s ease;
            width: auto;
            margin: 0;
        }
        button:hover {
            background-color: #9855d5;
        }
        #description {
            max-width: 80%;
            margin-bottom: 20px;
            font-size: 16px;
            line-height: 1.6;
            text-align: center;
            color: #c4c4c4;
        }
        .btn {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            flex-grow: 0;
            padding: 0 20px;
            margin-top: 50px;
        }

        @media (max-width: 768px) {
            form {
                flex-direction: column; /* Stack elements vertically */
                align-items: center; /* Center-align the flex items */
            }
            .text-container, .output-container, .btn {
                width: 100%; /* Full width for better readability */
            }
            .text-container {
                order: 1; /* Ensure input is first */
            }
            .btn {
                order: 2; /* Ensure button is second */
                margin-top: 20px; /* Provide some space above the button */
                margin-bottom: 20px; /* Provide space below the button before output */
            }
            .output-container {
                order: 3; /* Ensure output is last */
            }
            textarea, #output {
                height: auto; /* Adjust height to be more suitable for mobile */
                max-height: 300px; /* Set a max height for easier interaction */
            }
            #description {
                font-size: 14px; /* Slightly reduce font size for space efficiency */
                padding: 0 10px; /* Reduce padding to fit more text */
            }
            h1 {
                font-size: 20px; /* Reduce heading size for small screens */
            }
        }
    </style>

    <script>
    document.addEventListener('DOMContentLoaded', function() {
        var submitButton = document.querySelector('form button[type="submit"]');
        var errorContainer = document.getElementById('error-message');

        submitButton.addEventListener('click', function() {
            // errorContainer.style.display = 'none'; // Hide the container
            errorContainer.innerHTML = ''; // Clear any existing error message
        });
    });
</script>
</head>
<body>
    <div id="container" hx-ext="response-targets">
        <h1>Repeating-Key XOR Cracker</h1>
        <p id="description"> This tool is designed to decipher data encoded in
        Base64 or Hex that has been secured using a repeating-key XOR cipher.
        By methodically exploring various key lengths, employing transposition
        and frequency analysis techniques, it aims to deduce both the key size
        and the key itself. Once the key has been identified, the tool proceeds
        to decrypt the encoded information with the discovered key.</p>
        <div id="error-message" style="color: red; z-index: 100;"></div>
        <form hx-post="/decrypt" hx-target="#output" hx-target-error="#error-message" hx-swap="innerHTML" id="inputForm">
            <div class="text-container">
                <label for="inputData">Input Data:</label>
                <textarea id="inputData" name="inputData" rows="20"></textarea>
            </div>
            <div class="btn">
                <button type="submit">Process</button>
            </div>
            <div class="output-container" id="scrollable-output-container">
                <div id="output">Process ciphertext to see output.</div>
            </div>
        </form>
    </div>
</body>
<script>
document.body.addEventListener('htmx:configRequest', function(event) {
    // Hide and clear the error message when a new request is being configured
    var errorContainer = document.getElementById('error-message');
    // errorContainer.style.display = 'none'; // Hide the container
    errorContainer.innerHTML = ''; // Clear any existing error message
});
</script>
</html>
`
