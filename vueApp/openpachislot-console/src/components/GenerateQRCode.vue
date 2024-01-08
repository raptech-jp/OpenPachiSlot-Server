<template>
    <div v-if="qrCodeData" class="flex flex-col items-center">
        <img :src="qrCodeData" class="m-auto" alt="QR Code">
        <p class="mt-4">Name: {{ name }}</p>
        <p class="mt-2">この画面を保存してください。QRコードの再発行はできません。</p>
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
