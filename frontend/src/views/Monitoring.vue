<script setup lang="ts">
import { ref } from "vue";
import { useWebSocket } from '@vueuse/core';
import LineChart from '@/components/LineChart.vue';

type DataPoint = {
    time: number;
    total: number;
    used: number;
};

const chartData = ref<DataPoint[]>([]);


function bytesToGigabytes(bytes) {
    return Number((bytes / (1024 ** 3)).toFixed(3));
}

// Устанавливаем WebSocket соединение
const { status, data, send, close } = useWebSocket('ws://localhost:8080/api/v1/websocket/metrics', {
    onMessage: (message) => {
        // Парсим данные из сообщения
        const parsedData = JSON.parse(data.value);
        const dataPoint: DataPoint = {
            time: parsedData.time,
            total: bytesToGigabytes(parsedData.mem.total),
            used: bytesToGigabytes(parsedData.mem.used),
        };

        chartData.value = [...chartData.value, dataPoint];
    }
});

</script>

<template>
    <h2>Monitoring</h2>
    <LineChart :initial-data="chartData" />
</template>

<style>
    h2 {
        margin-bottom: 20px;
    }
</style>