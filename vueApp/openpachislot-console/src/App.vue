<template>
  <header ref="header" class="container mx-auto p-4 md:px-8 user-select-none">
    <h1 class="text-center font-bold text-5xl tracking-tight bg-gradient-to-r from-green-500 via-blue-500 to-pink-500 bg-clip-text text-transparent max-w-2xl mx-auto">
      OpenPachiSlot
    </h1>
  </header>
  <div class="overflow-auto">
    <router-view></router-view>
  </div>
</template>