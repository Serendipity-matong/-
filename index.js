const form = document.getElementById('form');  
const container =document.querySelector('#container');
const signInButton =document.querySelector('#signIn');
const signUpButton =document.querySelector('#signUp');
const usernameInput = document.getElementById('username');  
const passwordInput = document.getElementById('password');  
const passwordInput1 = document.getElementById('password1');  
const errorMessage = document.getElementById('error-message');  
const togglePasswordBtn = document.querySelector('.toggle-password'); 
const togglePasswordBtn1 = document.querySelector('.toggle-password1'); 
let isLogin = true;            
signUpButton.addEventListener('click',()=>{
    isLogin = !isLogin; 
    container.classList.add('right-panel-active'); 
});
signInButton.addEventListener('click',()=>{
    isLogin = !isLogin;
    container.classList.remove('right-panel-active');
});
togglePasswordBtn.addEventListener('click', () => {  
    passwordInput.type = passwordInput.type === 'password' ? 'text' : 'password';
    togglePasswordBtn.textContent = passwordInput.type === 'password' ? '展开' : '隐藏';  
});  
const userData = JSON.parse(localStorage.getItem('users')) || {};
form.addEventListener('submit', (e) => {
    e.preventDefault();  
    const username = usernameInput.value.trim(); 
    const password = passwordInput.value.trim();
    if (!isLogin) {  
        const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;  
        if (!emailPattern.test(username)) {  
            showError('格式错误');  
            return;  
        }  
        if (userData[username]) {  
            showError('邮箱已被注册');  
            return;  
        }  
          
        userData[username] = password;  
        localStorage.setItem('users', JSON.stringify(userData));  
        Swal.fire('注册成功', 'success');  
    } else {  
        if (!userData[username]) {  
            showError('用户名错误');  
            return;  
        }  
          
        if (userData[username] !== password) {  
            showError('密码错误');  
            return;  
        }  
        Swal.fire('登陆成功', 'success');  
    }  
});  
  
function showError(message) {  
    errorMessage.textContent = message;  
    errorMessage.style.display = 'block';  
}