const store = new  Vuex.Store({
    state: {
        tasks: [],
        places: []
    },

    mutations: {
        addTask(state, task){
            state.tasks.push(task)
        },
        addPlace(state, place){
            state.places.push(place)
        },
    },

    getters: {
        /*
        replacePlaceName: (state) => (id) =>{
            return null
        },

        /*
        task: (state) => (id) =>{
            return 1
        },

        /*
        warningTask: (state, getters) => {
            if(state.tasks.length == 0) return null
            var values = state.tasks.map((task) => {
                if (task.limit_time == 0){
                    return {task_id: task.id, value:
                        1 * (1 + getters.placeImportance(task.task_id)/100) * (1 + task.importance/100)
                    }
                }else return 1
            })

            values.sort((t1, t2) => {
                if (t1.value > t2.value) return -1
                else return 1
            })

            return getters.task(values[0])
        },
        */

        taskList: (state) => {
            return state.tasks
        }
    }
})

let warningTask = new Vue({
    el: '#warning-task',
    computed: {
        warningTask: function(){
            return null
        }
    }
})

let newTask = new Vue({
    el: '#new-task',
    data: {
        content: "",
    },

    methods:{
        addTask: function(){
            if(this.content.trim() == "")
                return null

            axios.post('/task/add', {content: this.content}).
                then((response) => store.commit('addTask', response.data)).
                catch((err) => console.log(err))

            this.content = "";
        }
    }
})

let taskList = new Vue({
    el: '#task-list',
    computed:{
        tasks: function(){
            return store.state.tasks
        }
    },

    created: function() {
        axios.post('/task/all')
            .then((response) => {
                response.data.forEach((data) => {
                    store.commit('addTask', data)
                })
            })

        axios.post('/place/all').
            then((response) => {
                response.data.forEach((data) =>{
                    store.commit('addPlace', data)
                })
            })
    }
})
