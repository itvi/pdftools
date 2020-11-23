FilePond.registerPlugin(
    FilePondPluginFileValidateType,
    FilePondPluginMediaPreview
);
const inputElement = document.querySelector('input[type="file"]');
const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/mergepdf',
    instantUpload: false,
    allowFileTypeValidation: true,
    acceptedFileTypes: ["application/pdf"],
    allowProcess: false,  // disable upload(up-arrow) icon on file
    allowReorder: true,
    labelIdle: '拖放文件或点击',
    labelFileTypeNotAllowed: '格式错误！',
    fileValidateTypeLabelExpectedTypes: '是PDF格式的文件吗？'
});

const mergeBtn = document.getElementById('upload');
mergeBtn.addEventListener('click',function(){
    var files = pond.getFiles();
    if(files.length<=1){
        $.notify({
            icon: 'fa fa-info-circle',
            message: '请选择多个PDF文件进行合并！',
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

    myAjaxUpload(files,"merge");
});
