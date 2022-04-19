
<template>
    <div class=" overflow-hidden sm:rounded-md">
      <div class="px-4 py-5 sm:p-6">
        <div class="grid grid-cols-6 gap-6">
          <div class="col-span-6 sm:col-span-3">
            <label for="book_name" class="block text-sm font-medium text-gray-700">Title</label>
            <input v-model="title" type="text" name="book_name" id="book_name" autocomplete="given-name" class="mt-1 py-2 px-3 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md">
          </div>
          <div class="col-span-6 sm:col-span-3">
            <label for="first_name" class="block text-sm font-medium text-gray-700">Author name</label>
            <input v-model="authorName" type="text" name="first_name" id="first_name" autocomplete="given-name" class="mt-1 py-2 px-3 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md">
          </div>
        </div>
        <button v-on:click="$emit('addBook', {Title:title,AuthorName: authorName, Created: '2020-11-10'})" class="inline-flex justify-center mt-4 py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
          Save
        </button>
      </div>
    </div>
</template>
<script>
export default {
  props:{
    authors: Array
  },
  methods: {
    format(author) {
      return JSON.parse(author);
    },
  },
  data() {
    return {
      title:"",
      authorName:"",
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