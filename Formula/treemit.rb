class Treemit < Formula
  desc "A tool for managing tree structures"
  homepage "https://github.com/yourusername/treemit"
  version "0.1.0"

  on_macos do
    url "file://#{ENV['HOME']}/my_codes/Go/treemit/treemit-0.1.0.tar.gz"
    sha256 "680b85d988d9d27c385fcd064d737deaa10ba08509bcbcac96d9f60e016d39e1"
  end

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "treemit", "./cmd/main/treemit.go"
    bin.install "treemit"
  end

  test do
    system "#{bin}/treemit", "--version"
  end
end