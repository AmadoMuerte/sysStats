<script setup lang="ts">
import { ref, reactive, watch, defineProps } from 'vue';
import VueApexCharts from 'vue3-apexcharts';

type DataPoint = {
    time: number;
    total: number;
    used: number;
};

const props = defineProps({
    initialData: {
        type: Array as () => DataPoint[],
        required: true
    }
})

const chartData = reactive({
    series: [
        {
            name: 'Total Memory (GB)',
            data: [] as number[]
        },
        {
            name: 'Used Memory (GB)',
            data: [] as number[]
        },
    ],
    options: {
        chart: {
            type: 'line',
            animations: {
                enabled: true,
                speed: 800,
                dynamicAnimation: {
                    enabled: true,
                    speed: 350
                }
            },
            toolbar: {
                show: false
            },
            zoom: {
                enabled: false
            }
        },
        stroke: {
            curve: 'smooth'
        },
        dataLabels: {
            enabled: false
        },
        xaxis: {
            categories: [] as string[],
            labels: {
                show: false
            }
        }
    }
});

watch(() => props.initialData, (newData) => {
    // Обновляем данные графика
    chartData.series[0].data = newData.map(item => item.total);
    chartData.series[1].data = newData.map(item => item.used);
    
    // Обновляем категории по временным меткам
    chartData.options.xaxis.categories = newData.map(item => {
        const date = new Date(item.time * 1000);
        return date.toLocaleString(); 
    });
}, { immediate: true });
</script>

<template>
    <div class="memoryChart">
        <VueApexCharts :options="chartData.options" :series="chartData.series" />
    </div>
</template>

<style>
    .memoryChart {
        width: 30%;
        max-height: auto;
        color: #000;
        background: #ffffff;
        border-radius: 10px;
    }
</style>
