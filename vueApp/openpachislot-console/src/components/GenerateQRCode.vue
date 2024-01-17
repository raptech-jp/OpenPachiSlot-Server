<template>
    <div v-if="qrCodeData" class="absolute inset-0 flex flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0 gap-2">
        <img :src="qrCodeData" alt="QR Code">
        <p>Name: {{ name }}</p>
        <p>この画面を保存してください。<br>QRコードの再発行はできません。</p>
    </div>
</template>

<script>
import QRCode from 'qrcode';

export default {
    props: {
        uuid: String,
        name: String,
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
                this.$emit('qr-generated', true); // 親コンポーネントにイベントを送信
            } catch (error) {
                console.error('Error generating QR code: ', error);
            }
        }
    },
};
</script>
