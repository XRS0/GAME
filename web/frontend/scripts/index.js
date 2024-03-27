function login() {
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    fetch('/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Успешный вход:', data);
    })
    .catch((error) => {
        console.error('Ошибка при входе:', error);
    });
}
