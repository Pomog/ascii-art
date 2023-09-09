// scripts.js

// Fetch the JSON data from the server
fetch('/get-strings')
    .then(response => response.json())
    .then(data => {
        const stringList = document.getElementById('stringList');
        data.forEach(string => {
            const listItem = document.createElement('li');
            listItem.textContent = string;
            stringList.appendChild(listItem);
        });
    })
    .catch(error => console.error('Error:', error));
