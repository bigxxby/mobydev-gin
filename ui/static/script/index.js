
function change(selected) {
    // Hide all carValue divs
    const elements = document.querySelectorAll('.carValue');
    elements.forEach(element => {
        element.style.display = 'none';
    });

    // Display the selected carValue div
    const selectedElement = document.getElementById(`carValue${selected}`);
    if (selectedElement) {
        selectedElement.style.display = 'block';
    }
}


function check() {
    // Get all radio buttons with the name "carousel"
    const radios = document.querySelectorAll('input[name="carousel"]');

    // Loop through each radio button to find the selected one
    radios.forEach((radio) => {
        if (radio.checked) {
            change(radio.value);
            return;
        }
    });
}
