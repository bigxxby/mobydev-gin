document.addEventListener('DOMContentLoaded', () => {
    const toggleThemeButton = document.querySelector('.toggle-theme');
    const body = document.body;

    toggleThemeButton.addEventListener('click', () => {
        if (body.classList.contains('dark-mode')) {
            // Switch to light mode
            body.classList.remove('dark-mode');
            toggleThemeButton.textContent = '🌙';
        } else {
            // Switch to dark mode
            body.classList.add('dark-mode');
            toggleThemeButton.textContent = '☀️';
        }
    });
});
function checkForinput() {
    let email = document.getElementById('email').value;
    let password = document.getElementById('password').value;
    let button = document.getElementById('submitButton');

    if (email === '' || password === '') {
        button.style.backgroundColor = 'gray';
        button.disabled = true;
    } else {
        button.style.backgroundColor = 'blue';
        button.disabled = false;
    }
}

function submit() {
    let email = document.getElementById('email').value;
    let password = document.getElementById('password').value;

    const data = {
        email: email,
        password: password,
    };

    const loading = document.getElementById('loading');
    const submitButton = document.getElementById('submitButton');

    loading.style.display = 'block';  // показать индикатор загрузки
    submitButton.disabled = true;     // деактивировать кнопку отправки

    fetch("/api/signIn", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            loading.style.display = 'none'; // скрыть индикатор загрузки
            submitButton.disabled = false;  // активировать кнопку отправки

            return response.json().then(json => {
                if (!response.ok) {
                    throw new Error(`${response.status}: ${json.message}`);
                }
                // const token = response.headers.get('Authorization');
                let token = json.token
                if (!token) {
                  throw new Error('Authorization header is missing');
                }
                localStorage.setItem('token', token.replace('Bearer ', ''));
                return json;
            });
        })
        .then(data => {
            showPopupNotification(data.message)

            setTimeout(() => {
                redirect();
            }, 2000);
        })
        .catch(error => {
            showPopupNotification(error.message);
        });
}
function redirect() {
    window.location.href = '/'
}
function showPopupNotification(message) {
    const popupNotification = document.getElementById('popupNotification');
    const popupMessage = document.getElementById('popupMessage');

    popupMessage.textContent = message;
    popupNotification.style.display = 'block';

    setTimeout(() => {
        closePopupNotification();
    }, 2000);
}

function closePopupNotification() {
    const popupNotification = document.getElementById('popupNotification');
    popupNotification.style.display = 'none';
}
function togglePasswordVisibility() {
    const passwordInput = document.getElementById("password");
    const toggleButton = document.getElementById("togglePasswordBtn");

    if (passwordInput.type === "password") {
        passwordInput.type = "text";
        toggleButton.innerHTML = '(X)'
    } else {
        passwordInput.type = "password";
        toggleButton.innerHTML = '(O)'
    }
}


