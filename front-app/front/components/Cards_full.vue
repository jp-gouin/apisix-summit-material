<template>
  <div  class="animate"> 
      <div class="border bg-white rounded lg:rounded-none p-4 flex flex-col justify-between leading-normal">
        <div class="flex flex-row">
          <div class="rounded-full flex h-16 w-16 shadow-md border-gray-400 border">
            <img class="h-12 w-12 m-auto rounded-full" :src="icon">
          </div>
          <slot name="header" class="flex flex-col flex-wrap">
            <div class="flex flex-col mt-1 ml-4 items-baseline ">
              <div class="whitespace-no-wrap capitalize font-medium ">
                {{name}}
              </div>
            </div>
          </slot>
        </div>
        <slot name="header2" class="flex flex-col flex-wrap"></slot>
        <span class="whitespace-normal  overflow-auto mt-2 mb-2 break-words h-16 max-h-xs text-xs text-gray-600">{{description}}</span>
        <div class="flex m-auto w-full">
            <slot name="content" class="flex flex-col flex-wrap">
            
            </slot>
        </div>
      </div>
    </div>
</template>
<script>
export default {
    name: 'Cards',
    props: {
      icon: String,
      description: String,
      name: String,
    }
  }
</script>
<style scoped>
.max-h-xs{
  max-height: 4rem;
}
.hash path {
    fill: none;
    stroke-width: 30
  } 
.last path {
    fill: none;
    stroke-width: 30
  } 
.hash path {
      stroke-dasharray: 3000;
      stroke-dashoffset: 3000;
      animation: animatePath 2s 0.5s forwards;
    }
.last path {
      stroke-dasharray: 2000;
      stroke-dashoffset: 2000;
      animation: animatePath 2s 0.5s forwards;
    }
@keyframes animatePath {
  to {
    stroke-dashoffset: 0;
  }
}
</style>