const store = new  Vuex.Store({
    state: {
        tasks: [],
        places: []
    },

    mutations: {
        addTask(){},
        addPlaces(){},
        deleteDoneTask(){},
        deltePlace(){},
    },
})

var app = new Vue({
    el: '.new-task',
    data:{
        content : "",
        tasks : []
    },
    
    methods:{
        addTask: function(){
            if(this.content.trim() == "")
                return

            this.tasks.push(this.content)
            this.content = "";
        },
    }
})

var todolist = new Vue({
    el: 'list'
})
