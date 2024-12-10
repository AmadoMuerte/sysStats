<script setup lang="ts">
import { ref } from "vue";
import { useWebSocket } from '@vueuse/core';
import Charts from '@/components/Charts.vue';
import type { DataPoint } from "@/types";


const chartData = ref<DataPoint[]>([]);

function bytesToGigabytes(bytes: number) {
    return Number((bytes / (1024 ** 3)).toFixed(3));
}


const { status, data, send, close } = useWebSocket("ws://localhost:8080/api/v1/websocket/metrics", {
    onMessage: () => {
        const parsedData = JSON.parse(data.value);
        const net = {
            gbReceived: bytesToGigabytes(parsedData.net[0].bytesRecv),
            gbSent: bytesToGigabytes(parsedData.net[0].bytesSent)            
        }

        const dataPoint: DataPoint = {
            time: parsedData.time,
            total: bytesToGigabytes(parsedData.mem.total),
            used: bytesToGigabytes(parsedData.mem),
            cpuPercent: parsedData.cpu[0],
            net: net
        };

        if (chartData.value.length == 150) {
            chartData.value = chartData.value.slice(10);
        }
        chartData.value = [...chartData.value, dataPoint];
    }
});

</script>

<template>
    <h2>Monitoring</h2>
    <Charts :initial-data="chartData" />
</template>

<style>
    h2 {
        margin-bottom: 20px;
    }
</style>