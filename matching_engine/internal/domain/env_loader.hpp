#pragma once

#include <string>
#include <fstream>
#include <cstdlib>
#include <sstream>
#include <algorithm>

inline void LoadEnvFile(const std::string &path = ".env") {
  std::ifstream f(path);
  if (!f.is_open()) return;
  std::string line;
  while (std::getline(f, line)) {
    auto pos = line.find('#');
    if (pos == 0) continue;
    if (pos != std::string::npos) line.resize(pos);
    auto eq = line.find('=');
    if (eq == std::string::npos) continue;
    std::string key = line.substr(0, eq);
    std::string val = line.substr(eq + 1);
    auto trim = [](std::string &s) {
      s.erase(s.begin(), std::find_if(s.begin(), s.end(),
                                      [](unsigned char c) { return !std::isspace(c); }));
      s.erase(std::find_if(s.rbegin(), s.rend(),
                           [](unsigned char c) { return !std::isspace(c); })
                  .base(),
              s.end());
    };
    trim(key);
    trim(val);
    if (!key.empty() && !val.empty()) {
      ::setenv(key.c_str(), val.c_str(), /*overwrite=*/1);
    }
  }
}
