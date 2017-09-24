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

const newTask = Vue.component('todo-newtask', {
    template: '<div> new task </div>'
})


Vue.component('todo-tasks', {
    template: '<div> task list </div>'
})

Vue.component('todo-place', {
})


new Vue({
    el : '#app',
    delimiters : ['<{', '}>']

    components: {
        'newTask': newTask,
    }
})


