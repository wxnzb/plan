struct strbuf {
    int len;
    int alloc;
    char *buf;
}
//初始化 sb 结构体，容量为 alloc
void strbuf_init(struct strbuf *sb, size_t alloc)
{
    sb->buf=(char*)malloc(alloc*sizeof(char));
    if(sb->buf==NULL)
    {
        exit(EXIT_FAILURE);
    }
    sb->len=0;
    sb->alloc=alloc;
};
//将字符串填充到 sb 中，长度为 len, 容量为 alloc
void strbuf_attach(struct strbuf *sb, void *str, size_t len, size_t alloc)
{
   sb->len=len;
   sb->alloc=alloc;
   sb->buf=(char*)str;
}
//   释放 sb 结构体的内存（！！释放缓冲区，而不是整个结构体）
void strbuf_release(struct strbuf *sb)
{
    free(sb->buf);
    sb->buf=NULL;
}
//交换两个 strbuf
void strbuf_swap(struct strbuf *a, struct strbuf *b)
{
//    int temp;
//    temp=a->len;
//    a->len=b->len;
//    b->len=temp;
//    temp=a->alloc;
//    a->alloc=b->alloc;
//    b->alloc=temp;
//    char *tmp;
//    tmp=a->buf;
//    a->buf=b->buf;
//    b->buf=tmp;
//    a->buf=(char*)realloc(a->buf,(a->len+1)*sizeof(char));
//    b->buf=(char*)realloc(b->buf,(b->len+1)*sizeof(char));
//    if(a->buf==NULL||b->buf==NULL)
//    {
//     free(a->buf);
//     free(b->buf);
//     exit(EXIT_FAILURE);
//    }
    struct strbuf temp;
    temp=*a;
    *a=*b;
    *b=temp;
}
//将 sb 中的原始内存取出，并将 sz 设置为其 alloc 大小 (不太理解)
char *strbuf_detach(struct strbuf *sb, size_t *sz)
{
    *sz=sb->alloc;
    return sb->buf;
}
//比较两个 strbuf 的内存是否相同(包括存的东西)
int strbuf_cmp(const struct strbuf *first, const struct strbuf *second)
{
    // return first->len-second->len;
    if(first->len==second->len)
    return 0;
    else
    return 1;
}
//清空 sb
void strbuf_reset(struct strbuf *sb)
{
   for(int i=0;i<sb->len;i++)
   sb->buf[i]='\0';
   sb->len=0;
}
//确保在 len 之后 strbuf 中至少有 extra 个字节的空闲空间可用
void strbuf_grow(struct strbuf *sb, size_t extra)
{
    if(sb->alloc<sb->len+1+extra)
    sb->alloc=sb->len+1+extra;
    sb->buf=(char*)realloc(sb->buf,sb->alloc*sizeof(char));
    if(sb->buf==NULL)
    {
        exit(EXIT_FAILURE);
    }
}

