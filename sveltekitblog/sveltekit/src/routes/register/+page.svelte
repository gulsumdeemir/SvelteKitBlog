<script>
    let userName = '';
    let email = '';
    let password = '';
    let errorMessage = '';

    const register = async () => {
        errorMessage = '';

        if (!userName || !email || !password) {
            errorMessage = 'Kullanıcı adı, e-posta ve şifre gerekli!';
            return;
        }

        try {
            const response = await fetch('http://localhost:3001/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    userName: userName,
                    email: email,
                    userPassword: password  
                }),
            });

            if (response.ok) {
                alert('Kayıt başarılı! Giriş yapabilirsiniz.');
                window.location.href = '/login'; 
            } else {
                const errorData = await response.json();
                errorMessage = errorData.error || 'Kayıt sırasında bir hata oluştu';
            }
        } catch (error) {
            
            errorMessage = 'Kayıt sırasında bir hata oluştu: ' + (error instanceof Error ? error.message : 'Bilinmeyen hata');
        }
    };
</script>

<style>
    .register-container {
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
</style>

<div class="register-container">
    <h2>Kayıt Ol</h2>
    <input
        type="text"
        placeholder="Kullanıcı Adı"
        bind:value={userName}
        required
    />
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
    <button on:click={register}>Kayıt Ol</button>
    {#if errorMessage}
        <div class="error">{errorMessage}</div>
    {/if}
</div>
