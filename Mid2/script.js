document.addEventListener('DOMContentLoaded', function () {
    // Initialize button elements and saved answers
    const submitButton = document.querySelector('button[type="submit"]');
    const saveButton = document.querySelector('button[type="button"]');
    const clearButton = document.getElementById('clear');
    let savedAnswers = JSON.parse(localStorage.getItem('savedAnswers')) || {};

    // Function to clear results from prediction and saved answer sections
    function clearResults() {
        document.querySelector('.prediction').textContent = '';
        document.querySelector('.saved-answer').textContent = '';
    }

    // Function to save current answers to local storage
    function saveToLocalStorage() {
        localStorage.setItem('savedAnswers', JSON.stringify(savedAnswers));
    }

    // Event listener for the submit button
    submitButton.addEventListener('click', function (e) {
        e.preventDefault();
        clearResults();
        const name = document.getElementById('name').value;
        if (!name) {
            alert("Please enter a name.");
            return;
        }
        // Check if name is saved and display, otherwise fetch from API
        if (savedAnswers[name]) {
            document.querySelector('.saved-answer').textContent = `${name}: ${savedAnswers[name]}`;
        } else {
            fetch(`https://api.genderize.io/?name=${name}`)
                .then(response => response.json())
                .then(data => {
                    document.querySelector('.prediction').textContent = `Gender: ${data.gender}, Probability: ${data.probability}`;
                });
        }
    });

    // Event listener for the save button
    saveButton.addEventListener('click', function () {
        const name = document.getElementById('name').value;
        const selectedGenderRadio = document.querySelector('input[name="gender"]:checked');
        if (!name || !selectedGenderRadio) {
            alert("Please enter a name and select a gender.");
            return;
        }
        const selectedGender = selectedGenderRadio.value;
        savedAnswers[name] = selectedGender;
        saveToLocalStorage();
        alert("Data saved successfully.");
    });

    // Event listener for the clear button to clear all saved data
    clearButton.addEventListener('click', function () {
        clearResults();
        localStorage.clear();
        savedAnswers = {};
    });
});
