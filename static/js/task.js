const store = new Vuex.Store({
    state: {
        tasks : []
    },

    mutations :{
        addTask(state, data){
            state.tasks.push(data)
        },
    }
})

let todoNewTask = Vue.extend({
    template: '<div> new task </div>'
})



let todoTasks = Vue.extend({
    created: function() {
        axios.post('/task/all')
            .then(response => {
                response.data.forEach(task => store.commit('addTask', task))
            })
    },

    computed : {
        tasks() {
            return store.state.tasks
        }
    },

    template: temp,
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


