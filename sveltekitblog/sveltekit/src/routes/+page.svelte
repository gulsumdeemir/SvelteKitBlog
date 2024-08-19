<script>
    import { goto } from '$app/navigation';
</script>

<main>
    <h1>Blog Uygulaması</h1>
    <button on:click={() => goto('/login')}>Giriş Yap</button>
</main>
