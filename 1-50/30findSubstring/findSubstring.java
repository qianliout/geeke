class Sulution {

    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> res = new ArrayList<Integer>();
        int wordNum = words.length;
        if (wordNum == 0) {
            return res;
        }
        int wordLen = words[0].length();
        //HashMap1 存所有单词
        HashMap<String, Integer> allWords = new HashMap<String, Integer>();
        for (String w : words) {
            int value = allWords.getOrDefault(w, 0);
            allWords.put(w, value + 1);
        }
        //遍历所有子串
        for (int i = 0; i < s.length() - wordNum * wordLen + 1; i++) {
            //HashMap2 存当前扫描的字符串含有的单词
            HashMap<String, Integer> hasWords = new HashMap<String, Integer>();
            int num = 0;
            //判断该子串是否符合
            while (num < wordNum) {
                String word = s.substring(i + num * wordLen, i + (num + 1) * wordLen);
                //判断该单词在 HashMap1 中
                if (allWords.containsKey(word)) {
                    int value = hasWords.getOrDefault(word, 0);
                    hasWords.put(word, value + 1);
                    //判断当前单词的 value 和 HashMap1 中该单词的 value
                    if (hasWords.get(word) > allWords.get(word)) {
                        break;
                    }
                } else {
                    break;
                }
                num++;
            }
            //判断是不是所有的单词都符合条件
            if (num == wordNum) {
                res.add(i);
            }
        }
        return res;
    }
}

class Solution2 {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> rs = new ArrayList<>();
        if (words == null || words.length == 0) {
            return rs;
        }
        int wLen = words[0].length();
        int wTotalLen = words[0].length() * words.length;
        int sLen = s.length();
        int[] flagArray = new int[wLen];
        for (int i = 0; i < wLen; i++) {
            int flag = 0;
            for (String word : words) {
                flag += word.charAt(i);
            }
            flagArray[i] = flag;
        }

        BitSet bs = new BitSet(words.length);
        for (int i = 0; i <= sLen - wTotalLen; i++) {
            int j = 0;
            for (; j < wLen; j++) {
                int sumFlag = 0;
                for (int k = 0; k < words.length; k++) {
                    sumFlag += s.charAt(i + k * wLen + j);
                }
                if (sumFlag != flagArray[j]) {
                    break;
                }
            }
            if (j == wLen) {
                // 可能匹配
                bs.clear();
                for (int k = 0; k < words.length; k++) {
                    String w = s.substring(i + k * wLen, i + (k + 1) * wLen);
                    int n = 0;
                    for (; n < words.length; n++) {
                        if (!bs.get(n) && words[n].equals(w)) {
                            bs.set(n);
                            break;
                        }
                    }
                    if (n == words.length) {
                        break;
                    }
                }
                if (bs.cardinality() == words.length) {
                    rs.add(i);
                }

            }
        }

        return rs;
    }
}

