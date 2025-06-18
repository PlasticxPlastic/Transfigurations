<template>
  <div class="max-w-2xl mx-auto p-6">
    <h2 class="text-2xl font-bold text-center mb-6 text-blue-600">Transfiguration Year 1 Calculator</h2>
    
    <div class="bg-white rounded-xl shadow-lg p-6">
      <!-- Target Value Selector -->
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">Calculate Value</label>
        <select 
          v-model="targetValue" 
          class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          @change="resetResult"
        >
          <option value="T">Transfiguration Power (T)</option>
          <option value="W">Wand Power (W)</option>
          <option value="F">Force (F)</option>
          <option value="S">Matter Value (S)</option>
          <option value="g">Mass (g)</option>
        </select>
      </div>

      <form @submit.prevent="calculateValue" class="space-y-6">
        <!-- Wand Power (W) -->
        <div v-if="targetValue !== 'W'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Wand Power (W)</label>
          <input 
            type="number" 
            v-model="wandPower" 
            step="0.01"
            min="0.01"
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter wand power"
            required
          />
        </div>

        <!-- Force (F) -->
        <div v-if="targetValue !== 'F'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Force (F)</label>
          <input 
            type="number" 
            v-model="force" 
            step="0.01"
            min="0.01"
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter force"
            required
          />
        </div>

        <!-- Matter Value (S) -->
        <div v-if="targetValue !== 'S'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Matter Value (S)</label>
          <select 
            v-model="selectedMatter" 
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          >
            <option value="">Select matter type</option>
            <option value="ก๊าซ">ก๊าซ (4.0)</option>
            <option value="ของเหลว">ของเหลว (5.5)</option>
            <option value="ของแข็ง">ของแข็ง (8.5)</option>
            <option value="ของแข็งพิเศษ">ของแข็งพิเศษ(โลหะ/อัญมณี) (12.5)</option>
            <option value="ของผสม">ของผสม (17.5)</option>
          </select>
        </div>

        <!-- Mass (g) -->
        <div v-if="targetValue !== 'g'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Mass (g)</label>
          <input 
            type="number" 
            v-model="mass" 
            step="0.01"
            min="0.01"
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter mass"
            required
          />
        </div>

        <!-- Transfiguration Power (T) -->
        <div v-if="targetValue !== 'T'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Transfiguration Power (T)</label>
          <input 
            type="number" 
            v-model="transfigurationPower" 
            step="0.01"
            min="0.01"
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter transfiguration power"
            required
          />
        </div>

        <button 
          type="submit"
          class="w-full py-3 px-4 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-gradient-to-r from-blue-500 to-orange-500 hover:from-blue-600 hover:to-orange-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transform transition duration-150 ease-in-out hover:scale-[1.02]"
        >
          Calculate {{ targetValue.toUpperCase() }}
        </button>
      </form>

      <!-- Results -->
      <div v-if="result !== null" class="mt-6 p-4 bg-blue-50 rounded-lg">
        <h3 class="text-lg font-semibold text-blue-700">Result</h3>
        <p class="text-2xl font-bold text-blue-600 mt-2">{{ formatNumber(result) }}</p>
        <button
          @click="showSolution = true"
          class="mt-4 text-blue-600 hover:text-blue-800 text-sm font-medium flex items-center"
        >
          <span>View Detailed Solution</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>

      <!-- Error Message -->
      <div v-if="error" class="mt-4 p-4 bg-red-50 text-red-600 rounded-lg">
        {{ error }}
      </div>
    </div>

    <!-- Solution Modal -->
    <div v-if="showSolution" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div class="bg-white rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div class="p-6">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-xl font-bold text-gray-900">Detailed Solution</h3>
            <button @click="showSolution = false" class="text-gray-500 hover:text-gray-700">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <div class="space-y-4">
            <div class="p-4 bg-blue-50 rounded-lg">
              <h4 class="font-semibold text-blue-900 mb-2">Final Result</h4>
              <p class="text-2xl font-bold text-blue-600">{{ formatNumber(result) }}</p>
            </div>

            <div class="space-y-2">
              <h4 class="font-semibold text-gray-900">Step by step calculation:</h4>
              <div class="pl-4 space-y-3">
                <template v-if="targetValue === 'T'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">T = ((W*F)/(g*S)) * 100</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">T = (({{ formatNumber(wandPower) }} * {{ formatNumber(force) }})/({{ formatNumber(mass) }} * {{ formatNumber(matterValues[selectedMatter]) }})) * 100</p>
                  
                  <p class="text-gray-700">3. Calculate W*F:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">T = ({{ formatNumber(wandPower * force) }}/{{ formatNumber(mass * matterValues[selectedMatter]) }}) * 100</p>
                  
                  <p class="text-gray-700">4. Calculate the division:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">T = {{ formatNumber((wandPower * force)/(mass * matterValues[selectedMatter])) }} * 100</p>
                  
                  <p class="text-gray-700">5. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">T = {{ formatNumber(result) }}</p>
                </template>

                <template v-if="targetValue === 'W'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W = (T * g * S) / (100 * F)</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W = ({{ formatNumber(transfigurationPower) }} * {{ formatNumber(mass) }} * {{ formatNumber(matterValues[selectedMatter]) }}) / (100 * {{ formatNumber(force) }})</p>
                  
                  <p class="text-gray-700">3. Calculate numerator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W = {{ formatNumber(transfigurationPower * mass * matterValues[selectedMatter]) }} / {{ formatNumber(100 * force) }}</p>
                  
                  <p class="text-gray-700">4. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W = {{ formatNumber(result) }}</p>
                </template>

                <template v-if="targetValue === 'F'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">F = (T * g * S) / (100 * W)</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">F = ({{ formatNumber(transfigurationPower) }} * {{ formatNumber(mass) }} * {{ formatNumber(matterValues[selectedMatter]) }}) / (100 * {{ formatNumber(wandPower) }})</p>
                  
                  <p class="text-gray-700">3. Calculate numerator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">F = {{ formatNumber(transfigurationPower * mass * matterValues[selectedMatter]) }} / {{ formatNumber(100 * wandPower) }}</p>
                  
                  <p class="text-gray-700">4. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">F = {{ formatNumber(result) }}</p>
                </template>

                <template v-if="targetValue === 'S'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">S = (W * F) / (T * g / 100)</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">S = ({{ formatNumber(wandPower) }} * {{ formatNumber(force) }}) / ({{ formatNumber(transfigurationPower) }} * {{ formatNumber(mass) }} / 100)</p>
                  
                  <p class="text-gray-700">3. Calculate W*F:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">S = {{ formatNumber(wandPower * force) }} / ({{ formatNumber(transfigurationPower * mass) }} / 100)</p>
                  
                  <p class="text-gray-700">4. Calculate denominator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">S = {{ formatNumber(wandPower * force) }} / {{ formatNumber(transfigurationPower * mass / 100) }}</p>
                  
                  <p class="text-gray-700">5. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">S = {{ formatNumber(result) }}</p>
                  
                  <p v-if="identifiedMatter" class="mt-4 p-3 bg-green-50 text-green-700 rounded">
                    Identified Matter Type: {{ identifiedMatter }}
                  </p>
                </template>

                <template v-if="targetValue === 'g'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">g = (W * F) / (T * S / 100)</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">g = ({{ formatNumber(wandPower) }} * {{ formatNumber(force) }}) / ({{ formatNumber(transfigurationPower) }} * {{ formatNumber(matterValues[selectedMatter]) }} / 100)</p>
                  
                  <p class="text-gray-700">3. Calculate W*F:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">g = {{ formatNumber(wandPower * force) }} / ({{ formatNumber(transfigurationPower * matterValues[selectedMatter]) }} / 100)</p>
                  
                  <p class="text-gray-700">4. Calculate denominator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">g = {{ formatNumber(wandPower * force) }} / {{ formatNumber(transfigurationPower * matterValues[selectedMatter] / 100) }}</p>
                  
                  <p class="text-gray-700">5. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">g = {{ formatNumber(result) }}</p>
                </template>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

// Target value to calculate
const targetValue = ref('T')
const result = ref(null)
const formulaExplanation = ref('')
const error = ref('')
const identifiedMatter = ref(null)
const showSolution = ref(false)

// Input values
const wandPower = ref('')
const force = ref('')
const selectedMatter = ref('')
const mass = ref('')
const transfigurationPower = ref('')

// Matter value mapping
const matterValues = {
  'ก๊าซ': 4.0,
  'ของเหลว': 5.5,
  'ของแข็ง': 8.5,
  'ของแข็งพิเศษ': 12.5,
  'ของผสม': 17.5
}

// Format number to 2 decimal places
const formatNumber = (num) => {
  return Number(num).toFixed(2)
}

// Reset result when changing target value
const resetResult = () => {
  result.value = null
  error.value = ''
  identifiedMatter.value = null
  // Reset all input values
  wandPower.value = ''
  force.value = ''
  selectedMatter.value = ''
  mass.value = ''
  transfigurationPower.value = ''
}

// Identify matter type based on value
const identifyMatterType = (value) => {
  const tolerance = 0.1
  const matterRanges = {
    'ก๊าซ': { min: 3.90, max: 4.10 },
    'ของเหลว': { min: 5.40, max: 5.60 },
    'ของแข็ง': { min: 8.40, max: 8.60 },
    'ของแข็งพิเศษ(โลหะ/อัญมณี)': { min: 12.40, max: 12.60 },
    'ของผสม': { min: 17.40, max: 17.60 }
  }

  for (const [type, range] of Object.entries(matterRanges)) {
    if (value >= range.min - tolerance && value <= range.max + tolerance) {
      return type
    }
  }
  return null
}

const calculateValue = () => {
  try {
    error.value = ''
    identifiedMatter.value = null
    
    // Validate only the required inputs for each calculation type
    switch (targetValue.value) {
      case 'T':
        if (!wandPower.value || wandPower.value <= 0) throw new Error('Please enter a valid wand power')
        if (!force.value || force.value <= 0) throw new Error('Please enter a valid force')
        if (!selectedMatter.value) throw new Error('Please select a matter type')
        if (!mass.value || mass.value <= 0) throw new Error('Please enter a valid mass')
        break
      case 'W':
        if (!force.value || force.value <= 0) throw new Error('Please enter a valid force')
        if (!selectedMatter.value) throw new Error('Please select a matter type')
        if (!mass.value || mass.value <= 0) throw new Error('Please enter a valid mass')
        if (!transfigurationPower.value || transfigurationPower.value <= 0) throw new Error('Please enter a valid transfiguration power')
        break
      case 'F':
        if (!wandPower.value || wandPower.value <= 0) throw new Error('Please enter a valid wand power')
        if (!selectedMatter.value) throw new Error('Please select a matter type')
        if (!mass.value || mass.value <= 0) throw new Error('Please enter a valid mass')
        if (!transfigurationPower.value || transfigurationPower.value <= 0) throw new Error('Please enter a valid transfiguration power')
        break
      case 'S':
        if (!wandPower.value || wandPower.value <= 0) throw new Error('Please enter a valid wand power')
        if (!force.value || force.value <= 0) throw new Error('Please enter a valid force')
        if (!mass.value || mass.value <= 0) throw new Error('Please enter a valid mass')
        if (!transfigurationPower.value || transfigurationPower.value <= 0) throw new Error('Please enter a valid transfiguration power')
        break
      case 'g':
        if (!wandPower.value || wandPower.value <= 0) throw new Error('Please enter a valid wand power')
        if (!force.value || force.value <= 0) throw new Error('Please enter a valid force')
        if (!selectedMatter.value) throw new Error('Please select a matter type')
        if (!transfigurationPower.value || transfigurationPower.value <= 0) throw new Error('Please enter a valid transfiguration power')
        break
    }

    const W = parseFloat(wandPower.value)
    const F = parseFloat(force.value)
    const S = matterValues[selectedMatter.value] || 0
    const g = parseFloat(mass.value)
    const T = parseFloat(transfigurationPower.value)

    switch (targetValue.value) {
      case 'T':
        // T = ((W*F)/(g*S)) * 100
        if (g * S === 0) {
          throw new Error('Invalid input values: division by zero')
        }
        result.value = ((W * F) / (g * S)) * 100
        formulaExplanation.value = 'T = ((W*F)/(g*S)) * 100'
        break
      case 'W':
        // W = (T * g * S) / (100 * F)
        if (F === 0) {
          throw new Error('Invalid input values: division by zero')
        }
        result.value = (T * g * S) / (100 * F)
        formulaExplanation.value = 'W = (T * g * S) / (100 * F)'
        break
      case 'F':
        // F = (T * g * S) / (100 * W)
        if (W === 0) {
          throw new Error('Invalid input values: division by zero')
        }
        result.value = (T * g * S) / (100 * W)
        formulaExplanation.value = 'F = (T * g * S) / (100 * W)'
        break
      case 'S':
        // S = (W * F) / (T * g / 100)
        if (T * g === 0) {
          throw new Error('Invalid input values: division by zero')
        }
        result.value = (W * F) / (T * g / 100)
        identifiedMatter.value = identifyMatterType(result.value)
        formulaExplanation.value = 'S = (W * F) / (T * g / 100)'
        break
      case 'g':
        // g = (W * F) / (T * S / 100)
        if (T * S === 0) {
          throw new Error('Invalid input values: division by zero')
        }
        result.value = (W * F) / (T * S / 100)
        formulaExplanation.value = 'g = (W * F) / (T * S / 100)'
        break
    }
  } catch (err) {
    error.value = err.message
    result.value = null
  }
}
</script> 