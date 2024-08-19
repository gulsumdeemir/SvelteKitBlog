<script>
    import { writable } from 'svelte/store';
    import { goto } from '$app/navigation';

    const user = writable({ userName: '', eMail: '' });

    let email = '';
    let password = '';
    let errorMessage = '';

    const login = async () => {
        errorMessage = '';

        if (!email || !password) {
            errorMessage = 'E-posta ve şifre gerekli!';
            return;
        }

        const response = await fetch('http://localhost:3001/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: email,
                password: password
            }),
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('token', data.token);
            localStorage.setItem('userID', data.userID); 
            user.set({ userName: data.userName, eMail: data.email });
            alert('Giriş başarılı!');
            goto('/blog');
        } else {
            const errorData = await response.json();
            errorMessage = errorData.error || 'Giriş sırasında bir hata oluştu';
        }
    };
</script>

<style>
    .login-container {
        max-width: 360px;
        margin: 2rem auto;
        padding: 2rem;
        border-radius: 8px;
        box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
        background: #ffffff;
        border: 1px solid #e0e0e0;
    }

    h2 {
        text-align: center;
        color: #191d69;
        margin-bottom: 1.5rem;
    }

    input {
        width: 100%;
        padding: 0.75rem;
        margin-bottom: 1rem;
        border: 1px solid #dcdcdc;
        border-radius: 6px;
        box-sizing: border-box;
    }

    input:focus {
        border-color: #FF8000;
        outline: none;
        box-shadow: 0 0 5px rgba(255, 128, 0, 0.3);
    }

    button {
        width: 100%;
        padding: 0.75rem;
        background-color: #FF8000; 
        color: #ffffff;
        border: none;
        border-radius: 6px;
        cursor: pointer;
        font-size: 1rem;
        transition: background-color 0.3s;
    }

    button:hover {
        background-color: #e67300; 
    }

    .error {
        color: #FF8000; 
        margin-top: 1rem;
        text-align: center;
    }

    .register-link {
        text-align: center;
        margin-top: 1rem;
    }

    .register-link a {
        color: #FF8000;
        text-decoration: none;
    }

    .register-link a:hover {
        text-decoration: underline;
    }
</style>

<div class="login-container">
    <h2>Giriş Yap</h2>
    <input
        type="email"
        placeholder="E-posta"
        bind:value={email}
        required
    />
    <input
        type="password"
        placeholder="Şifre"
        bind:value={password}
        required
    />
    <button on:click={login}>Giriş Yap</button>
    {#if errorMessage}
        <div class="error">{errorMessage}</div>
    {/if}
    <div class="register-link">
        <p>Hesabın yok mu? <a href="/register">Kayıt Ol</a></p>
    </div>
</div>
