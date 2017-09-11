
var todo = new Vue({
    el: '#todo',
    components :{
        'add-task': addTask
    }
})

var addTask = ({
    data: {
        content: "",
    },
    methods:{
        postTask : function(){
            axios.post('/test')
        }
    },

    template:
        `<div id="addTask" class="input-group">
            <input v-model="content" class="form-control">
            <span class="input-group-btn">
		        <button type="button" v-on:click="postTask" class="btn btn-default">Add Task</button>
	        </span>
         </div>`,
})