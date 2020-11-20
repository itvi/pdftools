FilePond.registerPlugin(FilePondPluginImagePreview);
const inputElement = document.querySelector('input[type="file"]');

const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/upload',
    instantUpload: false, // 不自动上传；默认为自动。自动上传one by one ，每次一个，循环。
    labelIdle: '拖放文件或点击',
    labelFileProcessing: '转换中',
    labelFileProcessingComplete: '完成',
    labelTapToCancel: '点击取消',
    labelTapToUndo: ''
});

// finished processing a file
// click button in imagePreview 
pond.on('processfile', (error, file) => {
    if (error) {
        console.log('Oh no');
        return;
    }
    console.log('File processed', file); // success

    var fileName = JSON.parse(file.serverId); // from server

    var a = document.createElement("a");
    var url = "http://"+window.location.host+"/download/"+fileName;
    //var linkText = document.createTextNode("single file");
    //a.appendChild(linkText);
    //a.title = "this is a title of a"
    //a.classList.add("btn", "btn-primary")
    a.setAttribute('href', url);
    a.setAttribute('download', fileName);
    document.body.appendChild(a);
    a.click();
});

// manual upload
const uploadBtn = document.getElementById('upload');
uploadBtn.addEventListener("click", function () {
    var files = pond.getFiles();
    if(files.length==0){
        $.notify({
            icon: 'fa fa-info-circle',
	        // title: 'Bootstrap notify',
	        message: '请选择图片!',
	        url: 'http://www.baidu.com',
	        target: '_blank'
        },{
            type: "info",
            //allow_dismiss: true,
            placement:{
                from:"top",
                align: "center"
            },
            animate: {
				enter: "animate__animated animate__fadeInDown",
				exit: "animate__animated animate__fadeOutUp"
			}
        });
        return;
    }

    myAjaxUpload(files,"img2pdf");
})

