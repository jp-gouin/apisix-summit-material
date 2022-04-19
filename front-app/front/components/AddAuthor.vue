
<template>
    <div class=" overflow-hidden sm:rounded-md">
      <div class="px-4 py-5 sm:p-6">
        <div class="grid grid-cols-12 gap-12">
          <div class="col-span-4 sm:col-span-3">
            <label for="first_name" class="block text-sm font-medium text-gray-700">First name</label>
            <input v-model="name" type="text" name="first_name" id="first_name" autocomplete="given-name" class="mt-1 py-2 px-3  focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md">
          </div>

          <div class="col-span-4 sm:col-span-3">
            <label for="last_name" class="block text-sm font-medium text-gray-700">Last name</label>
            <input v-model="lastname" type="text" name="last_name" id="last_name" autocomplete="family-name" class="mt-1 py-2 px-3  focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md">
          </div>
           <div class="col-span-4 sm:col-span-3">
            <label for="last_name" class="block text-sm font-medium text-gray-700">Phone Number</label>
            <input v-model="number" type="text" name="last_name" id="last_name" autocomplete="family-name" class="mt-1 py-2 px-3  focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md">
          </div>
        </div>
        <button v-on:click="$emit('addAuthor', {name:name,lastname: lastname,number: number})" class="inline-flex justify-center mt-4 py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
          Save
        </button>
      </div>
    </div>
</template>
<script>
export default {
  data() {
    return {
      name:"",
      lastname:"",
      number: "",
      columns: [
        {
          label: "Name",
          jsonId: ""
        },
        {
          label: "Lastname",
          jsonId: ""
        }
      ]
    };
  },
}
</script>