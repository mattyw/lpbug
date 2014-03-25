function! s:Lpbug(word)
    let pattern = '\d\+'
    let l:wordUC = matchstr(a:word, pattern)
    let x = system("lpbug ".l:wordUC)
endfunction

command! LPBug call s:Lpbug(expand("<cWORD>"))
