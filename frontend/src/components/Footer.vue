
<script setup>

const copyright = new Date().getFullYear();
</script>

<template>
    <div class="footer">
        <div class="container">
            <div class="footer__inner">
                <div class="footer__copy">
                    &copy; {{ copyright }} AmadoMuerte
                </div>
                <div class="footer__links">
                    <a target="_blank" href="https://github.com/AmadoMuerte">Github</a>
                    <a target="_blank" href="https://t.me/AmadoMuerte">Telegram</a>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
    .footer__inner {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .footer__links {
        display: flex;
        gap: 20px;
    }
</style>