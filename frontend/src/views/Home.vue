<template>
  <div class="max-w-4xl mx-auto">
    <div class="bg-white rounded-xl shadow-lg p-8">
      <h1 class="text-3xl font-bold text-center text-blue-600 mb-6">Welcome to Transfiguration</h1>
      <p class="text-gray-600 text-center text-lg">
        A system for calculating and learning transfiguration magic
      </p>
      <div class="mt-8 grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h2 class="text-xl font-semibold text-blue-700 mb-3">Wand Power Calculator</h2>
          <p class="text-gray-600 mb-4">Calculate your wand's power based on material, core, and length</p>
          <router-link 
            to="/wand-power"
            class="inline-block bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
          >
            Go to Calculator
          </router-link>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <h2 class="text-xl font-semibold text-orange-700 mb-3">Learn Transfiguration</h2>
          <p class="text-gray-600 mb-4">Study transfiguration techniques from basics to advanced</p>
          <router-link 
            to="/transfiguration-year1"
            class="inline-block bg-orange-600 text-white px-4 py-2 rounded-lg hover:bg-orange-700 transition-colors"
          >
            Start Learning
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template> 