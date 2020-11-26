FilePond.registerPlugin(
    FilePondPluginFileValidateType,
    FilePondPluginMediaPreview
);
const inputElement = document.querySelector('input[type="file"]');
const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/splitpdf',
    instantUpload: false,
    allowFileTypeValidation: true,
    acceptedFileTypes: ["application/pdf"],
    allowProcess: false,  // disable upload(up-arrow) icon on file
    labelIdle: '拖放文件或点击',
    labelFileTypeNotAllowed: '格式错误！',
    fileValidateTypeLabelExpectedTypes: '是PDF格式的文件吗？'
});

const splitBtn = document.getElementById('upload');
splitBtn.addEventListener('click',function(){
    var files = pond.getFiles();
    if(files.length==0){
        notify('请选择PDF格式的文件！');
        return;
    }

    myAjaxUpload(files,"split");
});
