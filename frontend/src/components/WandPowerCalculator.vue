<template>
  <div class="max-w-2xl mx-auto p-6">
    <h2 class="text-2xl font-bold text-center mb-6 text-blue-600">Wand Power Calculator</h2>
    
    <div class="bg-white rounded-xl shadow-lg p-6">
      <!-- Target Value Selector -->
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">Calculate Value</label>
        <select 
          v-model="targetValue" 
          class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          @change="resetResult"
        >
          <option value="W">Wand Power (W)</option>
          <option value="a">Wood Type Value (a)</option>
          <option value="b">Core Value (b)</option>
          <option value="L">Wand Length (L)</option>
          <option value="f">Flexibility (f)</option>
        </select>
      </div>

      <form @submit.prevent="calculateValue" class="space-y-6">
        <!-- Wood Type (a) -->
        <div v-if="targetValue !== 'a'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Wood Type (a)</label>
          <select 
            v-model="selectedWoodType" 
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          >
            <option value="">Select wood type</option>
            <optgroup label="Hard Woods">
              <option v-for="(wood, name) in hardWoods" :key="name" :value="name">{{ name }}</option>
            </optgroup>
            <optgroup label="Soft Woods">
              <option v-for="(wood, name) in softWoods" :key="name" :value="name">{{ name }}</option>
            </optgroup>
          </select>
        </div>

        <!-- Core Type (b) -->
        <div v-if="targetValue !== 'b'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Wand Core (b)</label>
          <select 
            v-model="selectedCore" 
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          >
            <option value="">Select core type</option>
            <optgroup label="Damage/Death - Small">
              <option v-for="(core, name) in damageSmallCores" :key="name" :value="name">{{ name }}</option>
            </optgroup>
            <optgroup label="Damage/Death - Large">
              <option v-for="(core, name) in damageLargeCores" :key="name" :value="name">{{ name }}</option>
            </optgroup>
            <optgroup label="Control/Travel - Small">
              <option v-for="(core, name) in controlSmallCores" :key="name" :value="name">{{ name }}</option>
            </optgroup>
            <optgroup label="Control/Travel - Large">
              <option v-for="(core, name) in controlLargeCores" :key="name" :value="name">{{ name }}</option>
            </optgroup>
          </select>
        </div>

        <!-- Flexibility (f) -->
        <div v-if="targetValue !== 'f'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Flexibility (f)</label>
          <select 
            v-model="flexibility" 
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          >
            <option value="0.10">Very Flexible</option>
            <option value="0.05">Medium Flexible</option>
            <option value="0.00">Not Flexible</option>
          </select>
        </div>

        <!-- Wand Length (L) -->
        <div v-if="targetValue !== 'L'" class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">Wand Length (L)</label>
          <input 
            type="number" 
            v-model="length" 
            step="0.01"
            min="0.01"
            class="block w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter wand length"
            required
          />
        </div>

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
        <div v-if="identifiedWoodType" class="mt-2 p-2 bg-green-50 text-green-700 rounded">
          Identified Wood Type: {{ identifiedWoodType }}
        </div>
        <div v-if="identifiedCoreType" class="mt-2 p-2 bg-green-50 text-green-700 rounded">
          Identified Core Type: {{ identifiedCoreType }}
        </div>
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
                <template v-if="targetValue === 'W'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = [a² + b²]/L - f</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = [{{ formatNumber(woodTypeValue) }}² + {{ formatNumber(coreTypeValue) }}²]/{{ formatNumber(length) }} - {{ formatNumber(flexibility) }}</p>
                  
                  <p class="text-gray-700">3. Calculate squares:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = [{{ formatNumber(Math.pow(woodTypeValue, 2)) }} + {{ formatNumber(Math.pow(coreTypeValue, 2)) }}]/{{ formatNumber(length) }} - {{ formatNumber(flexibility) }}</p>
                  
                  <p class="text-gray-700">4. Calculate sum in brackets:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = {{ formatNumber(Math.pow(woodTypeValue, 2) + Math.pow(coreTypeValue, 2)) }}/{{ formatNumber(length) }} - {{ formatNumber(flexibility) }}</p>
                  
                  <p class="text-gray-700">5. Calculate division:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = {{ formatNumber((Math.pow(woodTypeValue, 2) + Math.pow(coreTypeValue, 2))/length) }} - {{ formatNumber(flexibility) }}</p>
                  
                  <p class="text-gray-700">6. Calculate final W²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W² = {{ formatNumber(result * result) }}</p>
                  
                  <p class="text-gray-700">7. Take square root for final W:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">W = {{ formatNumber(result) }}</p>
                </template>

                <template v-if="targetValue === 'a'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = L(W² + f) - b²</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(length) }}({{ formatNumber(wandPower) }}² + {{ formatNumber(flexibility) }}) - {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²</p>
                  
                  <p class="text-gray-700">3. Calculate W²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(length) }}({{ formatNumber(wandPower * wandPower) }} + {{ formatNumber(flexibility) }}) - {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²</p>
                  
                  <p class="text-gray-700">4. Calculate W² + f:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(length) }}({{ formatNumber(wandPower * wandPower + flexibility) }}) - {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²</p>
                  
                  <p class="text-gray-700">5. Calculate L(W² + f):</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(length * (wandPower * wandPower + flexibility)) }} - {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²</p>
                  
                  <p class="text-gray-700">6. Calculate b²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(length * (wandPower * wandPower + flexibility)) }} - {{ formatNumber(Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }}</p>
                  
                  <p class="text-gray-700">7. Calculate final a²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a² = {{ formatNumber(result * result) }}</p>
                  
                  <p class="text-gray-700">8. Take square root for final a:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">a = {{ formatNumber(result) }}</p>
                  
                  <p v-if="identifiedWoodType" class="mt-4 p-3 bg-green-50 text-green-700 rounded">
                    Identified Wood Type: {{ identifiedWoodType }}
                  </p>
                </template>

                <template v-if="targetValue === 'b'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = L(W² + f) - a²</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(length) }}({{ formatNumber(wandPower) }}² + {{ formatNumber(flexibility) }}) - {{ formatNumber(woodTypeValue) }}²</p>
                  
                  <p class="text-gray-700">3. Calculate W²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(length) }}({{ formatNumber(Math.pow(wandPower, 2)) }} + {{ formatNumber(flexibility) }}) - {{ formatNumber(woodTypeValue) }}²</p>
                  
                  <p class="text-gray-700">4. Calculate W² + f:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(length) }}({{ formatNumber(Math.pow(wandPower, 2) + parseFloat(flexibility)) }}) - {{ formatNumber(woodTypeValue) }}²</p>
                  
                  <p class="text-gray-700">5. Calculate L(W² + f):</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(length * (Math.pow(wandPower, 2) + parseFloat(flexibility))) }} - {{ formatNumber(woodTypeValue) }}²</p>
                  
                  <p class="text-gray-700">6. Calculate a²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(length * (Math.pow(wandPower, 2) + parseFloat(flexibility))) }} - {{ formatNumber(Math.pow(woodTypeValue, 2)) }}</p>
                  
                  <p class="text-gray-700">7. Calculate final b²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b² = {{ formatNumber(result * result) }}</p>
                  
                  <p class="text-gray-700">8. Take square root for final b:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">b = {{ formatNumber(result) }}</p>
                  
                  <p v-if="identifiedCoreType" class="mt-4 p-3 bg-green-50 text-green-700 rounded">
                    Identified Core Type: {{ identifiedCoreType }}
                  </p>
                </template>

                <template v-if="targetValue === 'L'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = (a² + b²)/(W² + f)</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = ({{ formatNumber(hardWoods[selectedWoodType] || softWoods[selectedWoodType]) }}² + {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²)/({{ formatNumber(wandPower) }}² + {{ formatNumber(flexibility) }})</p>
                  
                  <p class="text-gray-700">3. Calculate a² and b²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = ({{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2)) }} + {{ formatNumber(Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }})/({{ formatNumber(wandPower) }}² + {{ formatNumber(flexibility) }})</p>
                  
                  <p class="text-gray-700">4. Calculate sum in numerator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = {{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }}/({{ formatNumber(wandPower) }}² + {{ formatNumber(flexibility) }})</p>
                  
                  <p class="text-gray-700">5. Calculate W²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = {{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }}/({{ formatNumber(wandPower * wandPower) }} + {{ formatNumber(flexibility) }})</p>
                  
                  <p class="text-gray-700">6. Calculate denominator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = {{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }}/{{ formatNumber(wandPower * wandPower + flexibility) }}</p>
                  
                  <p class="text-gray-700">7. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">L = {{ formatNumber(result) }}</p>
                </template>

                <template v-if="targetValue === 'f'">
                  <p class="text-gray-700">1. Start with the formula:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = (a² + b²)/L - W²</p>
                  
                  <p class="text-gray-700">2. Substitute the values:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = ({{ formatNumber(hardWoods[selectedWoodType] || softWoods[selectedWoodType]) }}² + {{ formatNumber(damageLargeCores[selectedCore] || damageSmallCores[selectedCore]) }}²)/{{ formatNumber(length) }} - {{ formatNumber(wandPower) }}²</p>
                  
                  <p class="text-gray-700">3. Calculate a² and b²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = ({{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2)) }} + {{ formatNumber(Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }})/{{ formatNumber(length) }} - {{ formatNumber(wandPower) }}²</p>
                  
                  <p class="text-gray-700">4. Calculate sum in numerator:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = {{ formatNumber(Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2)) }}/{{ formatNumber(length) }} - {{ formatNumber(wandPower) }}²</p>
                  
                  <p class="text-gray-700">5. Calculate division:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = {{ formatNumber((Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2))/length) }} - {{ formatNumber(wandPower) }}²</p>
                  
                  <p class="text-gray-700">6. Calculate W²:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = {{ formatNumber((Math.pow(hardWoods[selectedWoodType] || softWoods[selectedWoodType], 2) + Math.pow(damageLargeCores[selectedCore] || damageSmallCores[selectedCore], 2))/length) }} - {{ formatNumber(wandPower * wandPower) }}</p>
                  
                  <p class="text-gray-700">7. Final calculation:</p>
                  <p class="font-mono bg-gray-100 p-2 rounded">f = {{ formatNumber(result) }}</p>
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
const targetValue = ref('W')
const result = ref(null)
const formulaExplanation = ref('')
const error = ref('')
const identifiedWoodType = ref(null)
const identifiedCoreType = ref(null)
const showSolution = ref(false)

// Input values
const selectedWoodType = ref('')
const selectedCore = ref('')
const flexibility = ref('0.10')
const length = ref('')
const wandPower = ref('')

// Format number to 2 decimal places
const formatNumber = (num) => {
  return Number(num).toFixed(2)
}

// Add this function after the formatNumber function
const identifyCoreType = (value) => {
  const tolerance = 0.1
  const coreRanges = {
    'ทำร้าย/ทำให้ตาย - เล็ก': { min: 4.40, max: 5.10 },
    'ทำร้าย/ทำให้ตาย - ใหญ่': { min: 5.40, max: 6.10 },
    'เก็บเกี่ยว/ยินยอม - เล็ก': { min: 3.40, max: 4.10 },
    'เก็บเกี่ยว/ยินยอม - ใหญ่': { min: 2.40, max: 3.10 }
  }

  for (const [type, range] of Object.entries(coreRanges)) {
    if (value >= range.min - tolerance && value <= range.max + tolerance) {
      return type
    }
  }
  return null
}

// Add this function after the identifyCoreType function
const identifyWoodType = (value) => {
  const tolerance = 0.1
  const woodRanges = {
    'ไม้เนื้อแข็ง': { min: 3.50, max: 3.70 },
    'ไม้เนื้ออ่อน': { min: 2.80, max: 3.00 }
  }

  for (const [type, range] of Object.entries(woodRanges)) {
    if (value >= range.min - tolerance && value <= range.max + tolerance) {
      return type
    }
  }
  return null
}

// Reset result when changing target value
const resetResult = () => {
  result.value = null
  error.value = ''
  identifiedWoodType.value = null
  identifiedCoreType.value = null
  // Reset all input values
  selectedWoodType.value = ''
  selectedCore.value = ''
  flexibility.value = '0.10'
  length.value = ''
  wandPower.value = ''
}

// ข้อมูลวัสดุไม้
const materialMapping = {
  // ไม้เนื้ออ่อน
  "cedar": "ไม้เนื้ออ่อน",
  "cypress": "ไม้เนื้ออ่อน",
  "fir": "ไม้เนื้ออ่อน",
  "ivy": "ไม้เนื้ออ่อน",
  "larch": "ไม้เนื้ออ่อน",
  "pine": "ไม้เนื้ออ่อน",
  "prickly ash": "ไม้เนื้ออ่อน",
  "reed": "ไม้เนื้ออ่อน",
  "spruce": "ไม้เนื้ออ่อน",
  "tamarack": "ไม้เนื้ออ่อน",
  "vine": "ไม้เนื้ออ่อน",
  "yew": "ไม้เนื้ออ่อน",
  
  // ไม้เนื้อแข็ง
  "acacia": "ไม้เนื้อแข็ง",
  "applewood": "ไม้เนื้อแข็ง",
  "aspen": "ไม้เนื้อแข็ง",
  "birch": "ไม้เนื้อแข็ง",
  "black walnut": "ไม้เนื้อแข็ง",
  "blackthorn": "ไม้เนื้อแข็ง",
  "cherry": "ไม้เนื้อแข็ง",
  "chestnut": "ไม้เนื้อแข็ง",
  "dogwood": "ไม้เนื้อแข็ง",
  "ebony": "ไม้เนื้อแข็ง",
  "elm": "ไม้เนื้อแข็ง",
  "english oak": "ไม้เนื้อแข็ง",
  "hawthorn": "ไม้เนื้อแข็ง",
  "hazel": "ไม้เนื้อแข็ง",
  "holly": "ไม้เนื้อแข็ง",
  "hornbeam": "ไม้เนื้อแข็ง",
  "laurel": "ไม้เนื้อแข็ง",
  "mahogany": "ไม้เนื้อแข็ง",
  "maple": "ไม้เนื้อแข็ง",
  "pear": "ไม้เนื้อแข็ง",
  "poplar": "ไม้เนื้อแข็ง",
  "rosewood": "ไม้เนื้อแข็ง",
  "rowan swamp mayhaw": "ไม้เนื้อแข็ง",
  "silver lime": "ไม้เนื้อแข็ง",
  "sugar maple": "ไม้เนื้อแข็ง",
  "walnut": "ไม้เนื้อแข็ง",
  "willow": "ไม้เนื้อแข็ง",
}

// ข้อมูลแกนไม้กายสิทธิ์
const coreTypeMapping = {
  // ทำร้าย/ทำให้ตาย - เล็ก
  "กระดองปูไฟ": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "ขนโปเกรบิน": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "เขาแจ๊คคาโลป": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "เขี้ยวเทโบ": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "จงอยออเกอรีย์": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "ถุงน้ำโลบาลัก": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "ตาพิกซี่": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "ปีกแฟรี่": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "ผมวีร่า": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "หางซาลาแมนเดอร์": ["ทำร้าย/ทำให้ตาย", "เล็ก"],
  "กล้ามชิชเพอร์เฟิล": ["ทำร้าย/ทำให้ตาย", "เล็ก"],

  // ทำร้าย/ทำให้ตาย - ใหญ่
  "กระดูกเสือ": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "กะโหลกโทรลล์": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "ขนแมววอมปัส": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "เขาฮอร์นเซอร์เพ้นท์": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "กระดูกสันหลัง WRM": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "กระดูกเยติ": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "เอ็นหัวใจสนอลลี": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "เอ็นหัวใจมังกร": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],
  "ครีบฮิปโปแคมปัส": ["ทำร้าย/ทำให้ตาย", "ใหญ่"],

  // ยินยอม/เก็บเที่ยว - เล็ก
  "ก้านต้นดิตทานี": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "ขนฟินิกส์": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "ขนฮิปโปกริฟฟ์": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "เคราเดมิไกส์": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "ปะการังวิเศษ": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "เปลือกหอยวิเศษ": ["ยินยอม/เก็บเที่ยว", "เล็ก"],
  "หนามนาร์ล": ["ยินยอม/เก็บเที่ยว", "เล็ก"],

  // ยินยอม/เก็บเที่ยว - ใหญ่
  "ขนกริฟฟิน": ["ยินยอม/เก็บเที่ยว", "ใหญ่"],
  "ขนหางธันเดอร์เบิร์ด": ["ยินยอม/เก็บเที่ยว", "ใหญ่"],
  "ขนหางเธสตรอล": ["ยินยอม/เก็บเที่ยว", "ใหญ่"],
  "ผมเคลฟี": ["ยินยอม/เก็บเที่ยว", "ใหญ่"],
}

const coreValues = {
  "ทำร้าย/ทำให้ตาย - เล็ก": 4.5,
  "ทำร้าย/ทำให้ตาย - ใหญ่": 5.8,
  "ยินยอม/เก็บเที่ยว - เล็ก": 6.1,
  "ยินยอม/เก็บเที่ยว - ใหญ่": 4.7
}

// แยกไม้ตามประเภท
const hardWoods = computed(() => {
  return Object.fromEntries(
    Object.entries(materialMapping)
      .filter(([_, type]) => type === "ไม้เนื้อแข็ง")
  )
})

const softWoods = computed(() => {
  return Object.fromEntries(
    Object.entries(materialMapping)
      .filter(([_, type]) => type === "ไม้เนื้ออ่อน")
  )
})

// แยกแกนไม้กายสิทธิ์ตามประเภท
const damageSmallCores = computed(() => {
  return Object.fromEntries(
    Object.entries(coreTypeMapping)
      .filter(([_, [type, size]]) => type === "ทำร้าย/ทำให้ตาย" && size === "เล็ก")
  )
})

const damageLargeCores = computed(() => {
  return Object.fromEntries(
    Object.entries(coreTypeMapping)
      .filter(([_, [type, size]]) => type === "ทำร้าย/ทำให้ตาย" && size === "ใหญ่")
  )
})

const controlSmallCores = computed(() => {
  return Object.fromEntries(
    Object.entries(coreTypeMapping)
      .filter(([_, [type, size]]) => type === "ยินยอม/เก็บเที่ยว" && size === "เล็ก")
  )
})

const controlLargeCores = computed(() => {
  return Object.fromEntries(
    Object.entries(coreTypeMapping)
      .filter(([_, [type, size]]) => type === "ยินยอม/เก็บเที่ยว" && size === "ใหญ่")
  )
})

const woodTypeValue = computed(() => {
  if (!selectedWoodType.value) return 0
  const woodType = materialMapping[selectedWoodType.value]
  return woodType === "ไม้เนื้อแข็ง" ? 3.6 : 2.9
})

const coreTypeValue = computed(() => {
  if (!selectedCore.value) return 0
  const [type, size] = coreTypeMapping[selectedCore.value] || []
  if (!type || !size) return 0
  return coreValues[`${type} - ${size}`] || 0
})

const calculateValue = () => {
  error.value = ''
  result.value = null
  identifiedWoodType.value = null
  identifiedCoreType.value = null

  // Get the actual values from the mappings
  const woodValue = woodTypeValue.value
  const coreValue = coreTypeValue.value
  const wandLength = parseFloat(length.value) || 0
  const flexValue = parseFloat(flexibility.value) || 0
  const wandPowerValue = parseFloat(wandPower.value) || 0

  try {
    switch (targetValue.value) {
      case 'W':
        if (!woodValue || !coreValue || !wandLength) {
          error.value = 'Please fill in all required fields'
          return
        }
        const wSquared = (Math.pow(woodValue, 2) + Math.pow(coreValue, 2)) / wandLength - flexValue
        if (wSquared < 0) {
          error.value = 'Invalid input values: result would be imaginary'
          return
        }
        result.value = Math.sqrt(wSquared)
        break

      case 'a':
        if (!coreValue || !wandLength || !wandPowerValue) {
          error.value = 'Please fill in all required fields'
          return
        }
        const aSquared = wandLength * (Math.pow(wandPowerValue, 2) + flexValue) - Math.pow(coreValue, 2)
        if (aSquared < 0) {
          error.value = 'Invalid input values: result would be imaginary'
          return
        }
        result.value = Math.sqrt(aSquared)
        identifiedWoodType.value = identifyWoodType(result.value)
        break

      case 'b':
        if (!woodValue || !wandLength || !wandPowerValue) {
          error.value = 'Please fill in all required fields'
          return
        }
        const bSquared = wandLength * (Math.pow(wandPowerValue, 2) + flexValue) - Math.pow(woodValue, 2)
        if (bSquared < 0) {
          error.value = 'Invalid input values: result would be imaginary'
          return
        }
        result.value = Math.sqrt(bSquared)
        identifiedCoreType.value = identifyCoreType(result.value)
        break

      case 'L':
        if (!woodValue || !coreValue || !wandPowerValue) {
          error.value = 'Please fill in all required fields'
          return
        }
        const denominator = Math.pow(wandPowerValue, 2) + flexValue
        if (denominator === 0) {
          error.value = 'Invalid input values: division by zero'
          return
        }
        result.value = (Math.pow(woodValue, 2) + Math.pow(coreValue, 2)) / denominator
        break

      case 'f':
        if (!woodValue || !coreValue || !wandLength || !wandPowerValue) {
          error.value = 'Please fill in all required fields'
          return
        }
        result.value = (Math.pow(woodValue, 2) + Math.pow(coreValue, 2)) / wandLength - Math.pow(wandPowerValue, 2)
        break
    }
  } catch (e) {
    error.value = 'An error occurred during calculation'
    console.error(e)
  }
}
</script> 