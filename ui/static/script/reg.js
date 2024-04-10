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
validateInput()
        function validateInput() {
            let email = document.getElementById('email').value;
            let password = document.getElementById('password').value;
            let confirmPassword = document.getElementById('confirmPassword').value;
            
            const validationBlockEmail = document.getElementById('validationBlockEmail');
            const validationBlockPassword = document.getElementById('validationBlockPassword');
            const validationBlockPasswordC = document.getElementById('validationBlockPasswordC');
            validationBlockEmail.innerHTML = ''; // очищаем валидационный блок перед проверкой
            validationBlockPassword.innerHTML = ''; // очищаем валидационный блок перед проверкой
            validationBlockPasswordC.innerHTML = ''; // очищаем валидационный блок перед проверкой

            // Валидация email
            if (!email) {
                validationBlockEmail.innerHTML += '<p>Почта обязательна</p>';
            } else if (!validateEmail(email)) {
                validationBlockEmail.innerHTML += '<p>Неверный формат почты</p>';
            }

            // Валидация пароля
            if (!password) {
                validationBlockPassword.innerHTML += '<p>Пароль обязателен</p>';
            } else {
                if (!validatePasswordContainsUppercase(password)) {
                    validationBlockPassword.innerHTML += '<p>Хотя бы одна буква высшего регистра</p>';
                }
                if (!validatePasswordContainsLowercase(password)) {
                    validationBlockPassword.innerHTML += '<p>Хотя бы одна буква нижнего регистра</p>';
                }
                if (!validatePasswordContainsDigit(password)) {
                    validationBlockPassword.innerHTML += '<p>Хотя бы одна цифра</p>';
                }
                if (!validatePasswordContainsSpecialChar(password)) {
                    validationBlockPassword.innerHTML += '<p>Хотя бы один специальный символ</p>';
                }
                if (!validateLatin(password)) {
                    validationBlockPassword.innerHTML += '<p>Только латинские буквы</p>';
                }
                if (password.length < 8) {
                    validationBlockPassword.innerHTML += '<p>Хотя бы 8 символов</p>';
                }
            }

            if (password !== confirmPassword) {
                validationBlockPasswordC.innerHTML += '<p>Пароли не совпадают</p>';
            }
            if ((validationBlockEmail.innerHTML !== '') || (validationBlockPassword.innerHTML !== '') || (validationBlockPasswordC.innerHTML !== '')) {
                submitButton.disabled = true
                submitButton.style.backgroundColor = 'gray'
            } else {
                submitButton.disabled = false
                submitButton.style.backgroundColor = 'blue'
            }
        }
        function validateLatin(password) {
            var nonLatinPattern = /[^a-zA-Z\s\d!@#$%^&*()-=_+[\]{}|;':",.<>/?\\]/;
            return !nonLatinPattern.test(password);
}

        function validatePasswordContainsUppercase(password) {
            return /[A-Z]/.test(password);
        }

        function validatePasswordContainsLowercase(password) {
            return /[a-z]/.test(password);
        }

        function validatePasswordContainsDigit(password) {
            return /\d/.test(password);
        }

        function validatePasswordContainsSpecialChar(password) {
            return /[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/.test(password);
        }

        function validateEmail(email) {
            const re = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
            return re.test(String(email).toLowerCase());
        }

        function submit() {
            let email = document.getElementById('email').value;
            let password = document.getElementById('password').value;
            let confirmPassword = document.getElementById('confirmPassword').value;
            let role = document.getElementById('role').value;

            const data = {
                email: email,
                password: password,
                confirmPassword: confirmPassword,
                role: role
            };

            const loading = document.getElementById('loading');
            const submitButton = document.getElementById('submitButton');

            loading.style.display = 'block';  // показать индикатор загрузки
            submitButton.disabled = true;     // деактивировать кнопку отправки

            fetch("api/reg", {
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
                        return json;
                    });
                })
                .then(data => {
                    // alert(data.message || 'Успешная регистрация'); // выводим сообщение из ответа или стандартное сообщение
                    
                    showPopupNotification(data.message)
                    
                })
                .catch(error => {
                    showPopupNotification(error.message);
                });
        }
        function showPopupNotification(message) {
            const popupNotification = document.getElementById('popupNotification');
            const popupMessage = document.getElementById('popupMessage');

            popupMessage.textContent = message;
            popupNotification.style.display = 'block';

            setTimeout(() => {
                closePopupNotification();
            }, 4000); 
        }

        function closePopupNotification() {
            const popupNotification = document.getElementById('popupNotification');
            popupNotification.style.display = 'none';
        }



