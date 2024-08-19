
<script>
    import { onMount } from 'svelte';
    import { page } from '$app/stores'; 
    import Sidebar from '../components/Sidebar.svelte';

    let showSidebar = true;

    
    $: showSidebar = !['/login', '/register'].includes($page.url.pathname);
</script>

<div style="display: flex;">
    {#if showSidebar}
        <Sidebar />
    {/if}
    <div style="flex-grow: 1; padding: 1rem;">
        <slot /> 
    </div>
</div>

<style>
    
</style>
