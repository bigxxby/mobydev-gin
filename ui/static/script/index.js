window.onload = function () {
    getProfile();
};

function getProfile() {
    let sessionId = getSessionId();
    let element = document.getElementById('profileWindow')
    if (sessionId !== null) {
        element.style.display = 'flex'
    }else {
        return 
    }
    
    data = {
        sessionId: sessionId,
    }

    fetch("/get-profile", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
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
        })
        .catch(error => {
            showPopupNotification(error.message);
        });
}
function displayUser(user) {
    const profileData = document.querySelector('.profileData');


    // const userIdDiv = document.createElement('div');
    // userIdDiv.textContent = `ID: ${user.id}`;
    // profileWindow.appendChild(userIdDiv);

    const userEmailDiv = document.createElement('div');
    userEmailDiv.textContent = `${user.email}`;
    profileData.appendChild(userEmailDiv);

    if (user.name == null) {

        const userNameDiv = document.createElement('div');
        userNameDiv.textContent = `Добро пожаловать, Пользователь!`;

        profileData.appendChild(userNameDiv);
    }

}


function getSessionId() {
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i].trim();
        if (cookie.startsWith('session_id=')) {
            return cookie.substring(11);
        }
    }
    return null;
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
