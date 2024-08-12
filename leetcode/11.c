//反转字符串函数
void reverse(char* s, int start, int end) {
    while (start < end) {
        char temp = s[start];
        s[start++] = s[end];
        s[end--] = temp;
    }
}

char * reverseWords(char * s){
    // 1. 移除多余空格
    int len = strlen(s);
    int fast = 0, slow = 0;
    // 移除字符串之前的空格
    while (s[fast] == ' ') {
        fast++;
    }
    // 移除单词之间多余的空格
    while (fast < len - 1) {
        if (s[fast] == ' ' && s[fast + 1] == ' ') {
            fast++;
        } else {
            s[slow++] = s[fast++];
        }
    }
    // 移除字符串后面的空格
    if (s[fast] == ' ') {
        s[slow] = '\0';
    } else {
        s[slow++] = s[fast];
        s[slow] = '\0';
    }

    
    // 2. 反转整个字符串
    reverse(s, 0, slow - 1);

    
    // 3. 反转每一个单词
    for (int i = 0; i < slow; i++) {
        int j = i;
        while (j < slow && s[j] != ' ') {
            j++;
        }
        reverse(s, i, j - 1);
        i = j;
    }

    return s;
}

