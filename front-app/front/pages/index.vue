<template>
  <div class="m-auto content-center text-center h-auto">
    <div>
      <div class="inline-flex bg-gray-100/50" style="background-color: rgba(133,133,133,0.5);">
        <Logo class="inline-block" />
        <h2 class="title align-middle inline-block">Welcome to the bookshelf PoC</h2>
      </div>
    <div class="flex flex-col m-12">
      <div class="flex">
        <Cardbase class="flex-1 mr-1" :icon="iconBook" :description="descriptionBook" name="Library"  >
          <div slot="content" class="flex flex-col flex-wrap mt-2 w-full content-around">
            <Books v-bind:books="books" />
            <span>{{errBooks}}</span>
          </div>
        </Cardbase>
        <Cardbase  class="flex-1" :icon="iconAuthor" :description="descriptionAuthor" name="Author listing"  >
          <div slot="content" class="flex flex-col flex-wrap mt-2 w-full content-around">
            <Authors v-bind:authors="authors" />
            <span>{{errAuthors}}</span>
          </div>
        </Cardbase>
      </div>  
      <div class="flex mt-4">
         <Cardbase class="flex-1" :icon="iconBook" :description="descriptionAddBook" name="Add a book in the library"  >
          <div slot="content" class="flex flex-col flex-wrap mt-2 w-full content-around">
            <AddBook v-bind:authors="authors" @addBook="addBook" />
           </div> 
         </Cardbase>
          <Cardbase class="flex-1 mr-1" :icon="iconAuthor" :description="descriptionAddAuthor" name="Add an author to the listing"  >
        
            <div slot="content" class="flex flex-col flex-wrap mt-2 w-full content-around">
              <AddAuthor @addAuthor="addAuthor" />
            </div> 
         </Cardbase>
      </div>
      <div class="flex overflow-auto xl:h-72 mt-4">
       <Log :messages="messages"></Log>
      </div>
    </div> 
    </div>
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import AddAuthor from '~/components/AddAuthor.vue'
import Authors from '~/components/Authors.vue'
import Books from '~/components/Books.vue'
import AddBook from '~/components/AddBook.vue'
import Log from '~/components/Log.vue'
import Cardbase from '~/components/Cards_full.vue'
export default {
  components: {
    Logo,
    AddAuthor,
    AddBook,
    Authors,
    Books,
    Log,
    Cardbase
  },
  data() {
    return {
      config: "",
      books:[],
      authors: [],
      messages: [],
      errBooks: "",
      errAuthors: "",
      iconBook: "/livre-ouvert.png",
      iconAuthor: "/pngegg.png",
      descriptionBook: "This is the library , all books are listed here",
      descriptionAuthor: "This is the author section, all author's personnal info are listed here",
      descriptionAddAuthor: "Add an author by entering the name and last name of the author",
      descriptionAddBook: "Add a book by entering the name of the book and the author",
    }
  },
  methods:{
    addLog(action, result){
      this.messages.unshift({action: action, result: result})
    },
    addAuthor(author){
      console.log(author)
      this.$axios.post(this.config.postAuthorUri, author, {headers:{'Authorization': this.config.token}})
      .then(response =>{
         console.log(response.data)
         this.addLog("AddAuthor",response.data)
         this.getAuthors()
        })
        .catch(e => {
          console.log(e);
          this.addLog("AddAuthor",e)
        });
    },
    addBook(book){
      console.log(book)
      this.$axios.post(this.config.postBookUri, book, {headers:{'Authorization': this.config.token}})
      .then(response =>{
         console.log(response.data)
         this.addLog("addBook",response.data)
         this.getBooks()
        })
        .catch(e => {
          console.log(e);
           this.addLog("addBook",e)
        });
    },
    refresh(){
      this.authors = []
      this.books = []
      this.init()
    },
    init(){
      this.$axios
        .get(window.location.origin+'/api/config')
        //.get('http://localhost:3000/api/config')
        .then(response =>{
          this.config = response.data
          this.addLog("init",response.data)
          this.getAuthors()
          this.getBooks()
        })
        .catch(e => {
          console.log(e);
          this.addLog("init",e)
        });
    },
    getAuthors(){
     this.$axios
        //.get(this.config.authorsUri , {withCredentials: true, headers:{'Access-Control-Request-Headers': 'Content-Type'}})
        .get(this.config.authorsUri , {headers:{'Authorization': this.config.token}})
        .then(response =>{
          this.authors = response.data.authors
          console.log(this.authors)
          this.addLog("getAuthors",response.data)
        })
        .catch(e => {
          console.log(e);
          this.addLog("getAuthors",e)
          this.errAuthors = e
          this.message = e
        });
    },
    getBooks(){
      this.$axios
        .get(this.config.booksUri,{headers:{'Authorization': this.config.token}})
        .then(response =>{
          this.books = response.data.books
          console.log('books')
          console.log(this.books)
          this.addLog("getBooks",response.data)
        })
        .catch(e => {
          console.log(e);
          this.errBooks = e
          this.addLog("getBooks",e)
        });
    }
  },
  mouted() {
    console.log("init")
    this.init()
  },
  mounted: function () {
   console.log("init")
    this.init()
  },
  computed: {
      token () {
        console.log(this.$cookies.getAll())
      return this.$cookies.getAll();
    }
  }
}
</script>

<style>
/* Sample `apply` at-rules with Tailwind CSS
.container {
@apply min-h-screen flex justify-center items-center text-center mx-auto;
}
*/
.bg{
  background-repeat: no-repeat;
  background-size: cover;
}
.min-heigth{
  min-height: 15rem;
}
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family:
    'Quicksand',
    'Source Sans Pro',
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    'Helvetica Neue',
    Arial,
    sans-serif;
  display: block;
  font-weight: 300;
  font-size: 50px;
  color: #ffffff;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 24px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
