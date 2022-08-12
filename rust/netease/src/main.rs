extern crate core;

use std::{fs, io};
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, LineWriter, Write};
use std::path::Path;


fn main() -> io::Result<()> {
    let list_dir: &Path = Path::new("list");
    if list_dir.is_file() {
        panic!("{} should be a directory", list_dir
            .as_os_str()
            .to_os_string()
            .into_string()
            .expect("cannot to string"))
    }

    let mut song_map: HashMap<String, Vec<String>> = HashMap::new();

    for entry in fs::read_dir("list")? {
        let entry = entry?;
        let path = entry.path();
        if path.is_dir() { continue; }

        let file = File::open(path)?;
        let mut lists: Vec<String> = Vec::new();
        let lines = io::BufReader::new(file)
            .lines().map(|l| l.unwrap());
        for line in lines {
            lists.push(line);
        }

        song_map.insert(String::from(entry
            .file_name()
            .into_string()
            .expect("want string")), lists);
    }

    let out_dir: &Path = Path::new("split_list");
    fs::create_dir(out_dir)?;

    for (name, songs) in song_map {
        println!("list name {}", name);
        let p = out_dir.join(&name);
        if !p.exists() {
            fs::create_dir(p)?;
        }

        let mut split_songs: Vec<Vec<String>> = Vec::new();
        for x in 0..=songs.len() / 500 {
            split_songs.insert(x, Vec::new())
        }
        for (idx, song) in songs.iter().enumerate() {
            let cur = idx / 500;
            let x = split_songs.get_mut(cur).unwrap();
            x.push(song.clone());
        }

        for (idx, list) in split_songs.iter().enumerate() {
            let file = File::create(out_dir.join(&name).join(format!("{}_{}", &name, idx)))?;
            let mut file = LineWriter::new(file);
            for s in list {
                file.write_all(format!("{}\n", s).as_bytes())?;
            }
            file.flush()?;
            continue;
        }
    }


    Ok(())
}
