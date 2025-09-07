import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";

// https://vite.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@Components': path.resolve(__dirname, './src/Components'),
            '@Locales': path.resolve(__dirname, './src/Locales'),
            '@Pages': path.resolve(__dirname, './src/Pages'),
            '@Stores': path.resolve(__dirname, './src/Stores'),
        }
    }
})
