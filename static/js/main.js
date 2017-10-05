const store = new  Vuex.Store({
    state: {
        tasks: [],
        places: []
    },

    mutations: {
        addTask(state, task){
            state.tasks.push(task)
        },
        addPlaces(state, place){
            state.places.push(place)
        },
        deleteDoneTask(){},
        deletePlace(){},
    },
    getters: {
        warningTask: (state) => {
            /*
            state.tasks.forEach((task) => {
            })
            */
            return null
        }
    }
})

let warningTask = new Vue({
    el: '#warning-task',
    computed: {
        getWarningTask: function(){
            return store.getters.warningTask
        }
    }
})

let newTask = new Vue({
    el: '#new-task',
    data:{
        content : "",
    },
    
    methods:{
        addTask: function(){
            if(this.content.trim() == "")
                return null

            this.content = "";
        }
    }
})

let todolist = new Vue({
    el: '#task-list',
    created: function(){
        ///task/allにリクエスト送ってstoreに収納
        axios.post('/task/all')
            .then((response) => {
                response.data.forEach((data) => {
                    store.commit('addTask', data)
                })
            })
    },

    computed:{
        tasks: function(){
            return store.state.tasks
        }
    },

    methods: {
        editTask: function(){
            //modal windowを開いて編集する
        }
    }
})

