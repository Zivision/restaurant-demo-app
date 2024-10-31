<script setup lang="ts">
import { RouterView, RouterLink } from 'vue-router';
</script>

<template>
  <nav>
    <router-link to="/">Home</router-link>
    <router-link to="/about">About</router-link>
  </nav>
  <router-view />
</template>

