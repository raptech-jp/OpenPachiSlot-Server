<template>
    <div v-if="qrCodeData">
        <img :src="qrCodeData" alt="QR Code">
        <p>Name: {{ name }}</p>
    </div>
</template>

<script>
import QRCode from 'qrcode';

export default {
    props: {
        uuid: String,
        name: String
    },
    data() {
        return {
            qrCodeData: null,
        };
    },
    watch: {
        uuid: {
            handler(newVal) {
                this.generateQRCode(newVal);
            },
            immediate: true
        }
    },
    methods: {
        async generateQRCode(uuid) {
            try {
                this.qrCodeData = await QRCode.toDataURL(uuid);
            } catch (error) {
                console.error('Error generating QR code: ', error);
            }
        }
    },
};
</script>
