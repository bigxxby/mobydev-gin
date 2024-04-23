
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
                
                let token = json.token;
                console.log(token)
                if (!token) {
                    throw new Error('Authorization header is missing');
                }
                document.cookie = `jwtToken=${token}; path=/;`;
                return json;
            });
        })
        .then(data => {
            showPopupNotification(data.message);
    
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


