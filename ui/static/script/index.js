window.onload = function () {
    getProfile();
};

function getProfile() {
    let token = getTokenFromLocalStorage()
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
function getTokenFromLocalStorage() {
    const token = localStorage.getItem('token');
  
    if (token) {
      return token;
    } else {
      console.error('Token not found in localStorage');
      return null;
    }
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


function clickHandler() {
    let token   = getTokenFromLocalStorage();
    logout(token ); // Вызов функции logout с передачей sessionId
}
function logout() {
    localStorage.removeItem("token")
    showPopupNotification('Выход из аккаунта...')
    profileWindow.style.display = 'none'
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
