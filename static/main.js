var result = document.getElementById("result")
const baseURL = window.location.href

submitButton = document.getElementById("submit")
submitButton.onclick = function() {
    var xhr = new XMLHttpRequest();
    var fullURL = document.getElementById("fullURL");
    var shortURL = document.getElementById("shortURL");
    if (shortURL.value == "shortcut") {
        shortURL.value = ""
    }
    var data = JSON.stringify({ "fullURL": fullURL.value, "shortcut": shortURL.value }); 
    xhr.open('POST', baseURL, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status = 200) {
                var response = JSON.parse(xhr.responseText)
                if (response.Success == true) {
                    result.innerHTML = ('Here\'s your link: ' + '<a href="' + baseURL + response.Shortcut + '">' + baseURL + response.Shortcut + '</a>')
                } else {
                    result.innerHTML = ('Error: ' + shortURL.value + ' already exists!')
                }
                
            } else {
                console.log("The server returned an error.")
            }
        }
      };
    console.log("Sending data: " + data)
    xhr.send(data)
}   