<script setup>
import { useRoute } from "vue-router";

const route = useRoute();

const nav = [
    {
        name: "Home",
        icon: "pi pi-home",
        link: "/"
    },
    {
        name: "Monitoring",
        icon: "pi pi-chart-line",
        link: "/monitoring"
    },
    {
        name: "Auth",
        icon: "pi pi-user",
        link: "/auth"
    }
];
</script>

<template>
    <nav class="navigation">
        <RouterLink 
            v-for="item in nav" 
            :key="item.link"
            :to="item.link" 
            :class="['navigation__item', { 'navigation__item-active': route.path === item.link }]"
        >
            <i :class="item.icon"></i>
        </RouterLink>
    </nav>
</template>

<style>
.navigation {
    position: fixed;
    padding: 30px 18px;
    background: var(--grey-3);
    backdrop-filter: blur(10px);
    width: 60px;
    left: 0;
    height: 100vh;
    display: flex;
    align-items: center;
    flex-direction: column;
    gap: 20px;
}

.navigation__item {
    font-size: 23px;
    color: var(--color-primary);
    transition: .2s;
}

.navigation__item:hover {
    transform: scale(1.2);
    color: var(--color-primary-focus);
}

.navigation__item-active {
    color: var(--color-primary-focus);
}
</style>