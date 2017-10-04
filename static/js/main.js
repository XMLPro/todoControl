const store = new  Vuex.Store({
    state: {
        tasks: [],
        places: []
    },

    mutations: {
        addTask(state, task){
            state.tasks.push(task)
        },
        addPlaces(){},
        deleteDoneTask(){},
        deltePlace(){},
    },
    getters: {
        warningTask: (state) => {
            //ヤバイタスク計算
            //taskがないときはnullを返す
            return "gnwegoaw"
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
    },

    computed:{
        tasks: function(){
            //taskを返す
            return ['a', 'a']
        }
    },

    methods: {
        editTask: function(){
            //modal windowを開いて編集する
        }
    }
})

