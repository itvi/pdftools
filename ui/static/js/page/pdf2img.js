FilePond.registerPlugin(
    FilePondPluginFileValidateType,
    FilePondPluginMediaPreview
);

const inputElement = document.querySelector('input[type="file"]');
const pond = FilePond.create(inputElement);

pond.setOptions({
    // add additional data 
    server:{
        url: 'http://' + window.location.host, // http://ip:port
        process:{
            url: '/upload',
            method: 'POST',
            ondata:(formData)=>{
                formData.append('action','pdf2img');
                return formData;
            }
        }
    },
    instantUpload: false,
    allowReorder: true,
    allowFileTypeValidation: true,
    acceptedFileTypes: ["application/pdf"],
    allowProcess: false,
    labelIdle: '拖放文件或点击',
    labelFileTypeNotAllowed: '格式错误！',
    fileValidateTypeLabelExpectedTypes: '是PDF格式的文件吗？'
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
    var url = "http://" + window.location.host + "/download/" + fileName;
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
var spinner = document.getElementById('spinner');

uploadBtn.addEventListener("click", function() {
    var el =document.getElementsByClassName('filepond--file-status-main');
    for(var i = 0;i<el.length;i++){
        var msg = el[i].innerHTML;
        if (msg!=""){
            notify('请检查!');
            return;
        }
    }
    
    // validate type
    if(pond.status == 2){ // error
        return;
    }
    
    var files = pond.getFiles();
    if (files.length == 0) {
        notify('请选择PDF文件！');
        return;
    }

    var format = document.getElementById('format').value;
    if (format=="0"){
        notify("请选择要转换的图片格式!")
        return;
    }
    var checked = document.getElementById('pdf2oneimg').checked;

    var obj = {};
    obj.action = "pdf2img";
    obj.format = format;
    obj.pdf2oneimg=checked;
    
    spinner.style.display = "block";
    this.hidden = true;
    
    myAjaxUpload(files, obj);
});
