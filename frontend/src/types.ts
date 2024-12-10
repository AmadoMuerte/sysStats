export type DataPoint = {
    time: number;
    total: number;
    used: number;
    cpuPercent: number;
    net: {
        gbReceived: number;
        gbSent: number;
    };
};