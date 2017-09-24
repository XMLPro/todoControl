const store = new Vuex.Store({
    state: {
        tasks : []
    },

    mutations :{
        initTasks(state, data) {
            state.tasks = data
        }
    }
})

let todoNewTask = Vue.extend({
    
    template: '<div> new task </div>'
})



let todoTasks = Vue.extend({
    template: '<div> task list </div>'
})

/*
Vue.component('todo-place', {
})
*/


new Vue({
    el : '#app',
    delimiters : ['<{', '}>'],

    components: {
        'todo-newtask': todoNewTask,
        'todo-tasks': todoTasks
    }
})


