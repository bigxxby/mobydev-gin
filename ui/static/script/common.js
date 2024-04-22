window.onload = function () {
    getProfile();
};
function getToken() {
    const cookieName = "jwtToken=";
    const decodedCookie = decodeURIComponent(document.cookie);
    const cookieArray = decodedCookie.split(';');

    for (let i = 0; i < cookieArray.length; i++) {
        let cookie = cookieArray[i].trim();

        if (cookie.indexOf(cookieName) === 0) {
            return cookie.substring(cookieName.length, cookie.length);
        }
    }

    return null;
}
function displayUser(user) {
    const profileData = document.querySelector('.profileData');


    const userEmailDiv = document.createElement('div');
    userEmailDiv.textContent = `${user.email}`;
    profileData.appendChild(userEmailDiv);

    if (user.name == null) {

        const userNameDiv = document.createElement('div');

        userNameDiv.textContent = `Добро пожаловать, Пользователь!`;

        profileData.appendChild(userNameDiv);
    } else {
        const userNameDiv = document.createElement('div');

        userNameDiv.textContent = `Добро пожаловать, ` + user.name;

        profileData.appendChild(userNameDiv);

    }
    const logoutBtn = document.createElement('div');
    logoutBtn.textContent = `Выйти`;
    logoutBtn.className = 'logout'
    logoutBtn.onclick = clickHandler; // Присвоение обработчика кнопке
    profileData.appendChild(logoutBtn);
}

function getProfile() {
    let token = getToken()
    let element = document.getElementById('profileWindow')
    if (token !== null) {
        element.style.display = 'flex'
    } else {
        return
    }
    fetch(`api/profile`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`

        }
    })
        .then(response => {
            return response.json().then(json => {
                if (!response.ok) {
                    throw new Error(`${response.status}: ${json.message}`);
                }
                return json;
            });
        })
        .then(data => {
            displayUser(data);
            // console.log()
        })
        .catch(error => {
            showPopupNotification(error.message);
            localStorage.removeItem('token');
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
function logout(name) {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
    showPopupNotification('Выход из аккаунта...')
    profileWindow.style.display = 'none'
}
function clickHandler() {
    logout("jwtToken"); // Вызов функции logout с передачей sessionId
}
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