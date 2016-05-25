#ifndef CPPJIEBA_SEGMENTBASE_H
#define CPPJIEBA_SEGMENTBASE_H

#include "limonp/Logging.hpp"
#include "PreFilter.hpp"
#include <cassert>


namespace cppjieba {

const char* const SPECIAL_SEPARATORS = " \t\n，。";

using namespace limonp;

class SegmentBase {
 public:
  SegmentBase() {
    XCHECK(ResetSeparators(SPECIAL_SEPARATORS));
  }
  ~SegmentBase() {
  }

  bool ResetSeparators(const string& s) {
    symbols_.clear();
    RuneStrArray runes;
    if (!DecodeRunesInString(s, runes)) {
      XLOG(ERROR) << "decode " << s << " failed";
      return false;
    }
    for (size_t i = 0; i < runes.size(); i++) {
      if (!symbols_.insert(runes[i].rune).second) {
        XLOG(ERROR) << s.substr(runes[i].offset, runes[i].len) << " already exists";
        return false;
      }
    }
    return true;
  }
 protected:
  unordered_set<Rune> symbols_;
}; // class SegmentBase

} // cppjieba

#endif
