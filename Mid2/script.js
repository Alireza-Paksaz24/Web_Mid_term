document.getElementById('nameForm').addEventListener('submit', function(event){
    event.preventDefault();
    const name = document.getElementById('name').value;
    
    document.getElementById('result').textContent = '';
    document.getElementById('error').textContent = '';

    fetch(`https://api.genderize.io/?name=${name}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.gender) {
                document.getElementById('result').textContent = `Gender: ${data.gender} with Probability: ${data.probability}`;
            } else {
                throw new Error('No prediction available for this name');
            }
        })
        .catch(error => {
            document.getElementById('error').textContent = `Error: ${error.message}`;
        });
});

document.getElementById('saveButton').addEventListener('click', function() {
    const name = document.getElementById('name').value;
    const result = document.getElementById('result').textContent;
    if (result) {
        document.getElementById('savedAnswer').textContent = `Name: ${name}, ${result}`;
    }
