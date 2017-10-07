const store = new  Vuex.Store({
    state: {
        tasks: [],
        places: []
    },

    mutations: {
        addTask(state, task) {
            state.tasks.push(task)
        },
        addPlace(state, place){
            state.places.push(place)
        },
    },

    getters: {
        task: (state) => (id) =>{
            for(var i=0; i<state.tasks.length; i++){
                if(id == state.tasks[i].id) return state.tasks[i]
            }
        },


        placeImportance: (state) => (id) => {
            for(var i=0; i<state.places.length; i++){
                if(id == state.places[i].id) return state.places[i].importance
            }

            return null
        },

        warningTask: (state, getters) => {
            if(state.tasks.length == 0) return null
            var values = state.tasks.map((task) => {
                if (task.limit_time == 0){
                    return {task_id: task.id, value:
                        1 * (1 + getters.placeImportance(task.task_id)/100) * (1 + task.importance/100)
                    }
                }else {
                    return {
                        task_id: task.id, value:
                        1* (1 + (task.workload == 0? 1:task.workload)/(task.limit_time)) * (1 + getters.placeImportance(task.task_id)/100) * (1 + task.importance/100)
                    }
                }
            })

            values.sort((t1, t2) => {
                if (t1.value > t2.value) return -1
                else return 1
            })

            return getters.task(values[0].task_id)
        },

        taskList: (state) => {
            return state.tasks
        }
    }
})

let warningTask = new Vue({
    el: '#warning-task',
    computed: {
        warningTask: function(){
            return store.getters.warningTask
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
                then((response) => store.commit('addTask', response.data))

            this.content = "";
        }
    }
})

let focusTask = Vue.extend({
    template: '#modal-template',
    data: function() {
        return {
            active: false,
            id: 0,
            task: {}
        }
    },
    methods: {
        open: function(id)  {
            this.id = id
            this.task = store.getters.task(id);
            this.active = true
        },
        close: function() {
            this.active = false
        },
        commit: function () {

        }
    },

    mounted: function(){
        this.$nextTick(function() {
            taskList.$on('open-modal', this.open)
            taskList.$on('close-modal', this.close)
        }.bind(this))
    }
})

let taskList = new Vue({
    el: '#task-list',
    components: {
        'focus-task': focusTask
    },

    computed:{
        tasks: function(){
            return store.state.tasks
        }
    },

    methods:{
        openModal: function (id) {
            this.$emit('open-modal', id)
        },
        closeModal: function () {
            this.$emit('close-modal')
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
