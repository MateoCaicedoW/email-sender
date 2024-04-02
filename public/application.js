function openConfirm(el){
    const title = el.dataset.title
    const description = el.dataset.description
    const deleteText = el.dataset.delete
    const cancelText = el.dataset.cancel
    const bgConfirmButton = el.dataset.bgconfirmbutton

    return Swal.fire({
        title: title, 
        text: description, 
        showCancelButton: true, 
        showCloseButton: true, 
        customClass:{
                container: 'top-0 left-0 right-0 bottom-0', 
                title: 'text-lg font-medium leading-6 mb-1 inline-block flex', 
                cancelButton: 'bg-white mt-3 w-[48%] rounded-md border border-gray-300 shadow-sm px-4 py-2 text-base font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:text-sm', 
                confirmButton: bgConfirmButton+` w-[48%] rounded-md border border-transparent shadow-sm px-4 py-2 text-base font-medium text-white focus:outline-none sm:text-sm`
            }, 
        buttonsStyling: false,
        reverseButtons: true, 
        confirmButtonText: deleteText, 
        cancelButtonText: cancelText
    })
}

var auxBack = performance.getEntriesByType("navigation");
if (auxBack[0].type === "back_forward") {
    location.reload(true);
}
