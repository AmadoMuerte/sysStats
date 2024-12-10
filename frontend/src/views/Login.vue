<script setup>
import {ref} from "vue"

const props = defineProps(['switchAuth'])

const email = ref('')
const password = ref('')


const handleCLick = () => {
    if (email.value.trim() === '' || password.value.trim() === '') {
        return    
    }

    props.switchAuth();
    console.log(email.value, password.value)
}

</script>

<template>
    <div class="login">
        <div class="container">
            <div class="login__inner">
                <div class="loginForm">
                    <h2>Login</h2>
                    <div class="form__group">
                        <label for="email">Email</label>
                        <input v-model="email" name="email" type="email" placeholder="Email">
                    </div>
                    <div class="form__group">
                        <label for="password">Password</label>
                        <input v-model="password"  name="password" type="password" placeholder="Password">
                    </div>
                    <button @click="handleCLick">Login</button>
                </div>
            </div>
        </div>
    </div>
</template>


<style>
    .login {

    }
    .login__inner {
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .loginForm {
        border: 1px solid #fff;
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        padding: 20px 30px;
        background: var(--white);
    }

    .form__group {
        display: flex;
        flex-direction: column;
    }

    .loginForm label {
        color: var(--grey-3);
        margin-bottom: 5px;
        font-style: italic;
        font-size: 14px;
    }

    .loginForm button {
        border: 1px solid var(--grey-3);
        border-radius: 5px;
        padding: 10px 15px;
        background: var(--grey-4);
        color: #fff;
    }

    .loginForm input {
        padding: 10px 15px;
        border-radius: 5px;
        background: var(--white);
        border: 1px solid var(--grey-3);
    }

    .loginForm h2 {
        text-align: center;
        text-transform: uppercase;
        color: var(--grey-3);
    }
</style>